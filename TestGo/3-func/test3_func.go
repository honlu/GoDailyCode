/*
函数声明和调用
2023-2-1 by lu

Go的函数特点：
	支持不定变参，多个返回值，匿名函数和闭包；函数也是一种类型，一个函数可以赋值给变量
	不支持嵌套(即一个包有两个名字一样的函数),不支持重载和默认参数.
*/
package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

//返回多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

//返回多个返回值， 有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---- foo3 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//r1 r2 属于foo3的形参，  初始化默认的值是0
	//r1 r2 作用域空间 是foo3 整个函数体的{}空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---- foo4 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo4("foo4", 444)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
}
