package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
2022-9-22
question:
	实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
	另外一个从 channel 中读取数字并打印到标准输出。
	最终输出五个随机数。
answer:
	golang 协程和channel配合使用

想让面试官满意的答案还是有几点注意的地方：
	1. goroutine 在golang中是非阻塞的
	2. channel 无缓冲情况下，读写都是阻塞的，且可以用for循环来读取数据，当管道关闭后，for 退出。
	3. golang 中有专用的select case 语法从管道读取数据。
*/

func main() {
	ch := make(chan int)

	// wg用来等待程序来完成
	wg := sync.WaitGroup{} // 注意用{}初始化
	wg.Add(2)              // 计数加2，表示要让main等待两个goroutine

	// 声明一个匿名函数，并创建一个goroutine
	// write
	go func() {
		defer wg.Done() // 在此匿名函数退出时，调用Done来通知main函数此工作已经完成

		for i := 0; i < 5; i++ {
			ch <- rand.Intn(5) // 产生随机数,写入到chan中。ch为无缓冲通道，当读入一个值时会等待被读取，再进行下一步
		}
		close(ch) // 关闭ch，这里必须close。5个数写完之后，如果不close(ch)，read就会阻塞
	}()

	// read
	go func() {
		defer wg.Done()     // 在此匿名函数退出时，调用Done来通知main函数此工作已经完成
		for i := range ch { // 等待接收ch通道里的值。注意：遍历一个未关闭的channel会报错。因为发送端没有继续向通道发送数据，继续接收数据就会deadlock
			fmt.Println(i)
		}
	}()
	wg.Wait() // 在此阻塞，等待wg维护的计数为0
}
