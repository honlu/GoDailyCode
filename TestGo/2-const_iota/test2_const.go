/*
关于常量定义以及iota的特性
2023-2-1 by lu
introduce:
	const来定义常量。
	iota比较特殊，在每一个 const 关键字出现时被重置为 0，
		然后在下一个 const 出现之前，每出现一次 iota，
		其所代表的数字会自动增 1。可以被认为是一个可被编译器修改的常量，

	枚举：可以算看作包含一系列相关的常量。
	Go中可以通过const后跟一括号定义一组常量的定义来实现枚举。
*/
package main

import "fmt"

//const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota， 每行的iota都会累加1, 第一行的iota的默认值是0
	BEIJING  = 10 * iota //iota = 0
	SHANGHAI             //iota = 1
	SHENZHEN             //iota = 2
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d                      // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f                      // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
	i, k                      // iota = 4, i = iota * 2, k = iota * 3 , i = 8, k = 12
)

func main() {
	//常量(只读属性)
	const length int = 10

	fmt.Println("length = ", length)

	//length = 100 //常量是不允许修改的。

	fmt.Println("BEIJIGN = ", BEIJING)   // 0
	fmt.Println("SHANGHAI = ", SHANGHAI) // 100
	fmt.Println("SHENZHEN = ", SHENZHEN) // 200

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota 只能够配合const() 一起使用， iota只有在const进行累加效果。
	//var a int = iota
}
