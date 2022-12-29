package main

import (
	"fmt"
	"io"
	"strings"
)

/*
与string值相比，Builder值的优势其实主要体现在字符串拼接方面。
strings.Builder类型的值（以下简称Builder值）的优势有下面的三种：
	已存在的内容不可变，但可以拼接更多的内容；
	减少了内存分配和内容拷贝的次数；
	可将内容重置，可重用值。
限制：
	在已被真正使用后就不可再被复制；(我们只要调用了Builder值的拼接方法或扩容方法，就意味着开始真正使用它了。显而易见，这些方法都会改变其所属值中的内容容器的状态。再复制会panic)
	由于其内容不是完全不可变的，所以需要使用方自行解决操作冲突和并发安全问题
*/
func example1() { // 关于关于string.Builders类型的拼接和扩容
	// 示例1。
	var builder1 strings.Builder
	// 通过调用WriteXxx方法把新的内容拼接到已存在的内容的尾部（也就是右边）。会自动扩容
	builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String())
	fmt.Println()
	builder1.WriteByte(' ')
	builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
	builder1.Write([]byte{'\n', '\n'})
	builder1.WriteString("Do not copy a non-zero Builder.")
	fmt.Printf("The second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String())
	fmt.Println()

	// 示例2。
	fmt.Println("Grow the builder ...")
	// 手动扩容
	builder1.Grow(10)
	fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())
	fmt.Println()

	// 示例3。
	fmt.Println("Reset the builder ...")
	// 让Builder值重新回到零值状态，就像它从未被使用过那样。
	builder1.Reset()
	fmt.Printf("The third output(%d):\n%q\n", builder1.Len(), builder1.String())
}

func example2() { // 关于string.Builders类型的复制和panic
	// 示例1。
	var builder1 strings.Builder
	builder1.Grow(1)

	f1 := func(b strings.Builder) {
		//b.Grow(1) // 这里会引发panic。
	}
	f1(builder1)

	ch1 := make(chan strings.Builder, 1)
	ch1 <- builder1
	builder2 := <-ch1
	//builder2.Grow(1) // 这里会引发panic。
	_ = builder2

	builder3 := builder1
	//builder3.Grow(1) // 这里会引发panic。
	_ = builder3

	// 示例2。
	f2 := func(bp *strings.Builder) {
		(*bp).Grow(1) // 这里虽然不会引发panic，但不是并发安全的。
		builder4 := *bp
		//builder4.Grow(1) // 这里会引发panic。
		_ = builder4
	}
	f2(&builder1)

	builder1.Reset() // 重置之后就可以复制和使用
	builder5 := builder1
	builder5.Grow(1) // 这里不会引发panic。
}

func example3() { // 关于strings.Reader类型的高效读取。
	// Reader值实现高效读取的关键就在于它内部的已读计数。计数的值就代表着下一次读取的起始索引位置。它可以很容易地被计算出来。Reader值的Seek方法可以直接设定该值中的已读计数值。
	// 示例1。
	reader1 := strings.NewReader(
		"NewReader returns a new Reader reading from s. " +
			"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("The size of reader: %d\n", reader1.Size())
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len())) // 计算出的已读计数。已读计数也代表着下一次读取的起始索引位置。strings.Reader类型的Len()返回的是内部容器中未被读取部分的长度，而不是其中已存内容的总长度。

	buf1 := make([]byte, 47)
	n, _ := reader1.Read(buf1)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	fmt.Println()

	// 示例2。
	buf2 := make([]byte, 21)
	offset1 := int64(64)
	n, _ = reader1.ReadAt(buf2, offset1) //
	fmt.Printf("%d bytes were read. (call ReadAt, offset: %d)\n", n, offset1)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	fmt.Println()

	// 示例3。
	offset2 := int64(17)
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)
	readingIndex, _ := reader1.Seek(offset2, io.SeekCurrent) // seek方法可以直接设定Reader类型的已读计数值。
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex)
	fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)

	n, _ = reader1.Read(buf2)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
}

func main() {
	// example1()
	// example2()
	example3()
}
