/*
简单的情况下可以不使用任何库，直接处理 os.Args；其实 Golang 的标准库提供了 flag 包来处理命令行参数.
flag包：参考Golang_puzzlers中article2章代码.

一下是两种常用的定义命令行flag参数的方法：
- flag.Type(flag 名, 默认值, 帮助信息) *Type
- flag.TypeVar(Type 指针, flag 名, 默认值, 帮助信息)

Windows终端运行方式：
	1. 不传递命令行参数运行：go run .\flag.go
		这样name使用默认值everyone
	2. 传递命令行参数：go run .\flag.go -name=jack
		flag解析之后，name为jack

*/
package main

import (
	"flag"
	"fmt"
)

// 定义一个字符串类型的命令行参数变量，在 Init() 函数中对其初始化
// 因此，命令行参数对应变量的定义和初始化是可以分开的
var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse() // // 把用户传递的命令行参数解析为对应变量的值
	fmt.Printf("Hello, %s!\n", name)
	// flag.Args() 函数返回没有被解析的命令行参数
	// func NArg() 函数返回没有被解析的命令行参数的个数
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg()) // 较少用到
}
