/*
2022-9-23
question:
	字符串转为byte数组，会发生内存拷贝吗？
answer:
	会。严格来说，只要是发生类型强转都会发生内存拷贝。
	再问：
	频繁的内存拷贝操作听起来对性能不大友好。
	有没有什么办法可以在字符串转成切片的时候不用发生拷贝呢？
	再答：
	通过反射和unsafe包在底层结构转换，不发生内存的拷贝。
	字符串在go的底层结构式StringHeader,切片在Go的底层结构是SliceHeader.
	即把 StringHeader 的地址强转成 SliceHeader 就行。
	下面代码实现！

	具体是否有效可查看：
	    cnblogs.com/cheyunhua/p/15570988.html
		性能测试
*/
package main

import (
	"reflect"
	"unsafe"
)

func main() {
	a := "aaa"
	// unsafe.Pointer(&a) 获取变量a的地址
	// (*reflect.StringHeader)(unsafe.Pointer(&a)) 把字符串a转换成底层结构的形式
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	// (*[]byte)(unsafe.Pointer(&ssh)) 把ssh底层结构体转成byte切片的指针
	// 再通过 *转为指针指向的实际内容
	b := *(*[]byte)(unsafe.Pointer(&ssh))
}
