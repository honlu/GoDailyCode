package main

import "fmt"

/* 注意defer panic执行顺序 */
func A() {
	defer A1()
	defer A2()
	panic("panicA")
}

func A1() {
	fmt.Println("A1")
}

func A2() {
	defer B1()
	panic("panicA2")
}

func B1() {
	p := recover()
	fmt.Println("B1")
	fmt.Println(p)
}

func main() {
	A()
}
