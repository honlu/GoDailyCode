/*
40-41. io包的接口和工具
io包提供了很多接口。
好处：提高不同程序实体之间的互操作性。
在 Go 语言中，对接口的扩展是通过接口类型之间的嵌入来实现的，这也常被叫做接口的组合。
在io包中，io.Reader的扩展接口有下面几种:
	io.ReadWriter, io.ReadCloser, io.ReadWriteCloser, io.ReadSeeker, io.ReadWriteSeeker
io.Reader接口的实现类型，有下面几项：
	*io.LimitedReader, *io.SectionReader, *io.teeReader, *io.multiReader, *io.pipe, *io.PipeReader。
扩展接口其实就是，通过向自身嵌入其他接口来达到扩展那个接口的目的。
把没有嵌入其他接口并且只定义了一个方法的接口叫做简单接口。
有的接口有着众多的扩展接口和实现类型，我们可以称之为核心接口。
io包中的核心接口只有 3 个，它们是：io.Reader、io.Writer和io.Closer。
核心接口io.Writer。基于它的扩展接口除了有我们已知的io.ReadWriter、io.ReadWriteCloser和io.ReadWriteSeeker之外，
还有io.WriteCloser和io.WriteSeeker。
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func io1() {
	// 示例1。
	builder := new(strings.Builder)
	_ = interface{}(builder).(io.Writer)
	_ = interface{}(builder).(io.ByteWriter)
	_ = interface{}(builder).(fmt.Stringer)

	// 示例2。
	reader := strings.NewReader("")
	_ = interface{}(reader).(io.Reader)
	_ = interface{}(reader).(io.ReaderAt)
	_ = interface{}(reader).(io.ByteReader)
	_ = interface{}(reader).(io.RuneReader)
	_ = interface{}(reader).(io.Seeker)
	_ = interface{}(reader).(io.ByteScanner)
	_ = interface{}(reader).(io.RuneScanner)
	_ = interface{}(reader).(io.WriterTo)

	// 示例3。
	buffer := bytes.NewBuffer([]byte{})
	_ = interface{}(buffer).(io.Reader)
	_ = interface{}(buffer).(io.ByteReader)
	_ = interface{}(buffer).(io.RuneReader)
	_ = interface{}(buffer).(io.ByteScanner)
	_ = interface{}(buffer).(io.RuneScanner)
	_ = interface{}(buffer).(io.WriterTo)

	_ = interface{}(buffer).(io.Writer)
	_ = interface{}(buffer).(io.ByteWriter)
	_ = interface{}(buffer).(io.ReaderFrom)

	_ = interface{}(buffer).(fmt.Stringer)

	// 示例4。
	src := strings.NewReader(
		"CopyN copies n bytes (or until an error) from src to dst. " +
			"It returns the number of bytes copied and " +
			"the earliest error encountered while copying.")
	dst := new(strings.Builder)
	written, err := io.CopyN(dst, src, 58) // 想从src中拷贝前58个字节到dst那里。第一参数必须实现io.Writer接口，第二参数必须实现io.Reader接口
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Written(%d): %q\n", written, dst.String())
	}
}

func main() {

}
