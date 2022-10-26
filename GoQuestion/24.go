/*
2022-9-25
question：
	channel使用，以及close channel放在哪里合适呢？
answer:
	代码测试
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var ch1 = make(chan int) // 初始化channel,非nil
	wg.Add(1)
	go func() {
		temp := <-ch1     // 第一次接收。
		fmt.Println(temp) // 打印15
		temp = <-ch1      // 第二次接收。从一个已关闭的通道中接收数据，如果缓冲区为空，则返回改类型的零值
		fmt.Println(temp) // 打印0
		wg.Done()
	}()
	go func() {
		ch1 <- 15  // 第一次发送
		close(ch1) // 第二次直接关闭
	}()
	wg.Wait()

}
