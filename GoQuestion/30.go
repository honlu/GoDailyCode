/*
2022-9-27
question:
	有四个goroutine，编号为1、2、3、4。每秒钟会有⼀个goroutine打印出它⾃⼰的编号，
	要求你编写⼀个程序，让输出的编号总是按照1、2、3、4、1、2、3、4、……
	的顺序打印出来。
answer:
	这是一道经典的任务编排问题，使用channel.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func() {
		for {
			fmt.Println("1")
			time.Sleep(1 * time.Second)
			ch2 <- 1 // 发送，block
			<-ch1    // 接收，block
		}
	}()
	go func() {
		for {
			<-ch2 // 接收，block
			fmt.Println("2")
			time.Sleep(1 * time.Second)
			ch3 <- 1 // block
		}
	}()
	go func() {
		for {
			<-ch3 // block
			fmt.Println("3")
			time.Sleep(1 * time.Second)
			ch4 <- 1 // block

		}
	}()
	go func() {
		for {
			<-ch4 // block
			fmt.Println("4")
			time.Sleep(1 * time.Second)
			ch1 <- 1 // block

		}
	}()

	select {} // block
}
