/*
每秒执行一次proc并保证程序不退出.
实现：使用 time.Ticker 实现定时器的功能，使用 recover() 函数捕获 panic 错误。

一个函数栈对应一个defer栈，每次执行defer要知道压入什么一个函数栈对应一个defer栈，每次执行defer要知道压入什么

多以proc的执行一定要和defer recover在一个函数里面
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second) // 等待
			func() {
				defer func() { // 注意一个函数栈对应一个defer栈，每次执行defer要知道压入什么
					if err := recover(); err != nil {
						fmt.Println("捕获异常:", err)
					}
				}() // defer一定要在proc之前
				proc()
			}()
		}
	}()

	select {} // 阻塞main
}

func proc() {
	panic("ok")
}
