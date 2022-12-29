package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个[]byte的对象池，每个对象为一个[]byte
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	a := time.Now().Unix()
	fmt.Println(a)
	// 不使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := make([]byte, 1024) // 这里时for循环内部，频繁的创建[]byte数组，然后gc销毁[]byte数组，会对GC造成压力。
		_ = obj
	}
	b := time.Now().Unix()
	fmt.Println(b)
	// 使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := bytePool.Get().(*[]byte) //  从对象池拿 []byte 对象
		_ = obj                         // 操作对象，这里不重要
		bytePool.Put(obj)               // 操作结束后，放回对象池
	}
	c := time.Now().Unix()
	fmt.Println(c)
	fmt.Println("without pool ", b-a, "s")
	fmt.Println("with    pool ", c-b, "s")
}

/* run时禁用掉编译器优化，才会体现出有pool的优势
go run -gcflags="-l -N" pool.go
1671524760
1671524780
1671524793
without pool  20 s
with    pool  13 s

*/
