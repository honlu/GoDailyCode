/*
42-43.byfio包中的数据类型
bufio是“buffered I/O”的缩写。这个代码包中的程序实体实现的 I/O 操作都内置了缓冲区。
bufio包中的数据类型主要有：
	Reader；
	Scanner；
	Writer和ReadWriter。
bufio.Reader类型值中的缓冲区好处：
	可以在大多数的时候降低读取方法的执行时间。
	虽然，读取方法有时还要负责填充缓冲区，
	但从总体来看，读取方法的平均执行时间一般都会因此有大幅度的缩短。
bufio.Reader类型拥有很多用于读取数据的指针方法，这里面有 4 个方法
可以作为不同读取流程的代表，它们是：Peek、Read、ReadSlice和ReadBytes。
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func bufio1() {
	comment := "Package bufio implements buffered I/O. " +
		"It wraps an io.Reader or io.Writer object, " +
		"creating another object (Reader or Writer) that " +
		"also implements the interface but provides buffering and " +
		"some help for textual I/O."
	basicReader := strings.NewReader(comment) // strings.Reader类型
	fmt.Printf("The size of basic reader: %d\n", basicReader.Size())
	fmt.Println()

	// 示例1。
	fmt.Println("New a buffered reader ...")
	reader1 := bufio.NewReader(basicReader) //  bufio.Reader类型
	fmt.Printf("The default size of buffered reader: %d\n", reader1.Size())
	// 此时reader1的缓冲区还没有被填充。
	fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	fmt.Println()

	// 示例2。
	bytes, err := reader1.Peek(7)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Peeked contents(%d): %q\n", len(bytes), bytes)
	fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	fmt.Println()

	// 示例3。
	buf1 := make([]byte, 7)
	n, err := reader1.Read(buf1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Read contents(%d): %q\n", n, buf1)
	fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	fmt.Println()

	// 示例4。
	fmt.Printf("Reset the basic reader (size: %d) ...\n", len(comment))
	basicReader.Reset(comment)
	fmt.Printf("Reset the buffered reader (size: %d) ...\n", reader1.Size())
	reader1.Reset(basicReader)
	peekNum := len(comment) + 1
	fmt.Printf("Peek %d bytes ...\n", peekNum)
	bytes, err = reader1.Peek(peekNum)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("The number of peeked bytes: %d\n", len(bytes))
	fmt.Println()

	// 示例5。
	fmt.Printf("Reset the basic reader (size: %d) ...\n", len(comment))
	basicReader.Reset(comment)
	size := 300
	fmt.Printf("New a buffered reader with size %d ...\n", size)
	reader2 := bufio.NewReaderSize(basicReader, size)
	peekNum = size + 1
	fmt.Printf("Peek %d bytes ...\n", peekNum)
	bytes, err = reader2.Peek(peekNum)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("The number of peeked bytes: %d\n", len(bytes))
}

func main() {
	bufio1()
}
