package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup    // 等待组
	var ch chan int          // nil channel
	var ch1 = make(chan int) // 创建channel
	fmt.Println(ch, ch1)     // <nil> 0xc000086060
	wg.Add(1)                // 等待组的计数器+1
	go func() {
		//ch <- 15  // 如果给一个nil的channel发送数据会造成永久阻塞
		//<-ch  // 如果从一个nil的channel中接收数据也会造成永久阻塞
		ret := <-ch1
		fmt.Println(ret)
		ret = <-ch1 // 从一个已关闭的通道中接收数据，如果缓冲区中为空，则返回该类型的零值
		fmt.Println(ret)
		wg.Done() // 等待组的计数器-1
	}()
	go func() {
		//close(ch1)
		ch1 <- 15 // 给一个已关闭通道发送数据就会包panic错误
		close(ch1)
	}()
	wg.Wait() // 等待组的计数器不为0时阻塞
}
