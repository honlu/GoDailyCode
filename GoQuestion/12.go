/*
2022-9-22
question:
	完成代码，要求每秒钟调用一次proc并保证程序不退出
answer:
	考察了两个知识点：
		1、定时执行执行任务
		2、捕获 panic 错误
	要求每秒钟执行一次，首先想到的就是 time.Ticker对象，
	该函数可每秒钟往chan中放一个Time,正好符合我们的要求。

	在 golang 中捕获 panic 一般会用到 recover() 函数。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		t := time.NewTicker(time.Second * 1) // 周期性计时器
		for {
			select {
			case <-t.C: // 当1秒到达后，t.C会有值，然后接收就成功了！
				go func() {
					defer func() {
						if r := recover(); r != nil { // recover捕获panic后返回ok
							fmt.Println(r)
						}
					}()
					proc()
				}()

			}
		}

	}()

	select {} // 程序不能退出，使用select{}阻塞main
}

func proc() {
	panic("ok")
}
