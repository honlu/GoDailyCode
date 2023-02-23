/*
朋友给的代码：为什么会死锁？

*/
package main

import (
	"fmt"
	"time"
)

var chan1 chan int{}
var chan2 chan int{}

func main() {
	// chan1 = make(chan int)
	// chan2 = make(chan int)
	// chan1 <- 1 // 在这里会阻塞，导致后面goroutine没有创建，main死锁了
	for i := 0; i < 100; i++ {
		go func() {
			<-chan1
			fmt.Println("cat")
			chan2 <- 2
		}()
		go func() {
			<-chan2
			fmt.Println("dog")
			chan1 <- 1
		}()
	}
	chan1 <- 1
	time.Sleep(time.Second)
}
