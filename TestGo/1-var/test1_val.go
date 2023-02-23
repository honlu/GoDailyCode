/*
关于变量声明的四种方式
2023-2-1 by lu
introduce：
	var三种(不省略数据类型):
		1.声明一个变量不复制
		2.var边声明边赋值
	省略数据类型 3. var省去数据类型，赋值通过类型自动匹配数据类型
		4. 短变量声明:=

关于变量重声明，可重名变量：
	可重名变量：即不同代码块中的重名变量。
	注意：对于同一个代码块，声明重名的变量是无法通过编译的，用短变量声明对已有变量重声明除外。

- 变量重声明中的变量一定是在某一个代码块内的。
	注意，这里的“某一个代码块内”并不包含它的任何子代码块，
	否则就变成了“多个代码块之间”。而可重名变量指的正是在多个代码块之间由相同的标识符代表的变量。
- 变量重声明是对同一个变量的多次声明，这里的变量只有一个。
	而可重名变量中涉及的变量肯定是有多个的。
- 不论对变量重声明多少次，其类型必须始终一致，具体遵从它第一次被声明时给定的类型。
	而可重名变量之间不存在类似的限制，它们的类型可以是任意的。
- 如果可重名变量所在的代码块之间，存在直接或间接的嵌套关系，那么它们之间一定会存在“屏蔽”的现象。但是这种现象绝对不会在变量重声明的场景下出现。
*/

package main

import "fmt"

//声明全局变量 方法一、方法二、方法三是可以的
var gA int = 100
var gB = 200

//用方法四来声明全局变量
// := 只能够用在 函数体内来声明
//gC := 200

func main() {
	//方法一：声明一个变量 默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四：(常用的方法) 省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	// =====
	fmt.Println("gA = ", gA, ", gB = ", gB)
	//fmt.Println("gC = ", gC)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)
	var kk, ll = 100, "Aceld"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	//多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
