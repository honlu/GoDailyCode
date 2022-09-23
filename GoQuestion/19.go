/*
2022-9-23
question:
	谈谈golang的内存逃逸，什么情况下会发生内存逃逸？
answer:
	逃逸分析是编译器用来，判断一个变量是放在栈上还是堆上。
	Go程序变量会携带一组校验数据，用来证明它的整个声明周期是否在运行时完全可知。
	如果变量通过了校验，就被编译器分配到栈上，否则就发生逃逸，分配到堆上。

	变量逃逸的几种典型情况：
	1、在方法或函数中，返回局部变量指针
	（局部变量原本应该在栈中分配和回收，但是由于返回时被外部引用，
	因此生命周期大于栈，所以溢出逃逸，变量分配到堆中）
	2、申请变量内存所占极大
	（栈空间不足，所以发生逃逸）
	3、变量类型不确定，动态类型发生逃逸
	4、闭包引用对象逃逸

	通过go build/run -gcflage=-m xx.go查看逃逸的情况
*/
package main

import "fmt"

type A struct {
	s string
}

// 这是上面提到的 "在方法内把局部变量指针返回" 的情况
func foo(s string) *A {
	a := new(A) // 使用go build -gcflage=-m xx.go会发现这里发生逃逸
	a.s = s
	return a //返回局部变量a,在C语言中妥妥野指针，但在go则ok，但a会逃逸到堆
}
func main() {
	a := foo("hello")
	b := a.s + " world"
	c := b + "!"   // 会发现c发生逃逸
	fmt.Println(c) // 这是因为 fmt.Println() 的参数类型定义为 interface{}
	// 在 Go 语言中，空接口即 interface{} 可以表示任意的类型，如果函数参数为 interface{}，编译期间很难确定其参数的具体类型，也会发生逃逸。
}
