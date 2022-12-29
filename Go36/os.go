/*
44-45. 使用os包中的API
都是平台不相关的API:这些 API 基于（或者说抽象自）操作系统，为我们使用操作系统的功能提供高层次的支持，但是，它们并不依赖于具体的操作系统。
os包中的 API 主要可以帮助我们使用操作系统中的文件系统、权限系统、环境变量、系统进程以及系统信号。
获取os.File类型的指针值，os包有这几个函数：
	Create、NewFile、Open和OpenFile。
对File值的操作模式主要是只读模式、只写模式和读写模式，由常量os.O_RDONLY、os.O_WRONLY和os.O_RDWR代表
还有些额外的操作模式：
	os.O_APPEND：当向文件中写入内容时，把新内容追加到现有内容的后边。
	os.O_CREATE：当给定路径上的文件不存在时，创建一个新文件。
	os.O_EXCL：需要与os.O_CREATE一同使用，表示在给定的路径上不能有已存在的文件。
	os.O_SYNC：在打开的文件之上实施同步 I/O。它会保证读写的内容总会与硬盘上的数据保持同步。
	os.O_TRUNC：如果文件已存在，并且是常规的文件，那么就先清空其中已经存在的任何内容。
文件的访问权限：
os.OpenFile函数的第三个参数perm代表的是权限模式，其类型是os.FileMode。但实际上，os.FileMode类型能够代表的，可远不只权限模式，它还可以代表文件模式（也可以称之为文件种类）。

*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
)

// ioTypes 代表了io代码包中的所有接口的反射类型。
var ioTypes = []reflect.Type{
	reflect.TypeOf((*io.Reader)(nil)).Elem(),
	reflect.TypeOf((*io.Writer)(nil)).Elem(),
	reflect.TypeOf((*io.Closer)(nil)).Elem(),

	reflect.TypeOf((*io.ByteReader)(nil)).Elem(),
	reflect.TypeOf((*io.RuneReader)(nil)).Elem(),
	reflect.TypeOf((*io.ReaderAt)(nil)).Elem(),
	reflect.TypeOf((*io.Seeker)(nil)).Elem(),
	reflect.TypeOf((*io.WriterTo)(nil)).Elem(),
	reflect.TypeOf((*io.ByteWriter)(nil)).Elem(),
	reflect.TypeOf((*io.WriterAt)(nil)).Elem(),
	reflect.TypeOf((*io.ReaderFrom)(nil)).Elem(),

	reflect.TypeOf((*io.ByteScanner)(nil)).Elem(),
	reflect.TypeOf((*io.RuneScanner)(nil)).Elem(),
	reflect.TypeOf((*io.ReadSeeker)(nil)).Elem(),
	reflect.TypeOf((*io.ReadCloser)(nil)).Elem(),
	reflect.TypeOf((*io.WriteCloser)(nil)).Elem(),
	reflect.TypeOf((*io.WriteSeeker)(nil)).Elem(),
	reflect.TypeOf((*io.ReadWriter)(nil)).Elem(),
	reflect.TypeOf((*io.ReadWriteSeeker)(nil)).Elem(),
	reflect.TypeOf((*io.ReadWriteCloser)(nil)).Elem(),
}

func os1() {
	// 示例1。
	file1 := (*os.File)(nil)
	fileType := reflect.TypeOf(file1)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Type %T implements\n", file1)
	for _, t := range ioTypes {
		if fileType.Implements(t) {
			buf.WriteString(t.String())
			buf.WriteByte(',')
			buf.WriteByte('\n')
		}
	}
	output := buf.Bytes()
	output[len(output)-2] = '.'
	fmt.Printf("%s\n", output)

	// 示例2。
	fileName1 := "something1.txt"
	filePath1 := filepath.Join(os.TempDir(), fileName1)
	var paths []string
	paths = append(paths, filePath1)
	dir, _ := os.Getwd()
	paths = append(paths, filepath.Join(dir[:len(dir)-1], fileName1))
	for _, path := range paths {
		fmt.Printf("Create a file with path %s ...\n", path)
		_, err := os.Create(path)
		if err != nil {
			var underlyingErr string
			if _, ok := err.(*os.PathError); ok {
				underlyingErr = "(path error)"
			}
			fmt.Printf("error: %v %s\n", err, underlyingErr)
			continue
		}
		fmt.Println("The file has been created.")
	}
	fmt.Println()

	// 示例3。
	fmt.Println("New a file associated with stderr ...")
	file3 := os.NewFile(uintptr(syscall.Stderr), "/dev/stderr") // 第一个参数：代表文件描述符的以及uintptr类型的值，第二参数用于表示文件名的而字符串值。
	if file3 != nil {
		file3.WriteString(
			"The Go language program writes something to stderr.\n") // 向标准错误输出上写入一些内容
	}
	fmt.Println()

	// 示例4。
	fmt.Printf("Open a file with path %s ...\n", filePath1)
	file4, err := os.Open(filePath1) // 打开一个文件，并返回包装了该文件的File值，只能读不能写。
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println("Write something to the file ...")
	_, err = file4.WriteString("something")
	var underlyingErr string
	if _, ok := err.(*os.PathError); ok {
		underlyingErr = "(path error)"
	}
	fmt.Printf("error: %v %s\n", err, underlyingErr)
	fmt.Println()

	// 示例5。
	fmt.Printf("Open a file with path %s ...\n", filePath1)
	file5a, err := os.Open(filePath1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf(
		"Is there only one file descriptor for the same file in the same process? %v\n",
		file5a.Fd() == file4.Fd())
	file5b := os.NewFile(file5a.Fd(), filePath1)
	fmt.Printf("Can the same file descriptor represent the same file? %v\n",
		file5b.Name() == file5a.Name())
	fmt.Println()

	// 示例6。
	fmt.Printf("Reuse a file on path %s ...\n", filePath1)
	file6, err := os.OpenFile(filePath1, os.O_WRONLY|os.O_TRUNC, 0666) // os.OpenFile函数是os.Create函数和os.Open函数的底层支持，它最为灵活。
	// 有 3 个参数，分别名为name、flag和perm。name指代的就是文件的路径。flag参数指的则是需要施加在文件描述符之上的模式，我在前面提到的只读模式就是这里的一个可选项。
	// 参数perm代表的也是模式，它的类型是os.FileMode，此类型是一个基于uint32类型的再定义类型。
	// 把参数flag指代的模式叫做操作模式，而把参数perm指代的模式叫做权限模式。
	// 操作模式限定了操作文件的方式(只读os.O_RDONLY、只写os.O_WRONLY、读写模式os.O_RDWR)，而权限模式则可以控制文件的访问权限
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	contents := "something"
	fmt.Printf("Write %q to the file ...\n", contents)
	n, err := file6.WriteString(contents)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("The number of bytes written is %d.\n", n)
	}
}

func main() {
	os1()
}
