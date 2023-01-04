/*
2023-1-3
question:
	写出下面的代码输出内容
answer：（注意在一个函数中panic不会阻止defer后面语句或函数的执行。）
打印后
打印中
打印前
panic：触发异常

异常代码提示，等退出状态
*/
package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
