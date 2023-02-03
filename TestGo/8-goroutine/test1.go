/*
Goroutine的并发控制和通信。场景实现
一个典型场景就是主进程通知名下所有子goroutine优雅退出运行。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	running := true
	f := func() {
		for running { // 利用标志变量一直运行或停止
			fmt.Println("sub proc running...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("sub proc exit") // 利用goroutine的机制，执行完自身退出。
	}
	go f()
	go f()
	go f()
	time.Sleep(2 * time.Second)
	running = false
	time.Sleep(3 * time.Second)
	fmt.Println("main proc exit")
}

/*
全局变量的优势是：
	简单方便，不需要过多繁杂的操作，通过一个变量就可以控制所有子goroutine的开始和结束。
缺点是：
	功能有限，由于架构所致，该全局变量只能是多读一写，否则会出现数据同步问题。
	当然也可以通过给全局变量加锁来解决这个问题，但那就增加了复杂度。另外这种方式不适合用于子goroutine间的通信，因为全局变量可以传递的信息很小；
	还有就是主进程无法等待所有子goroutine退出。
	因为这种方式只能是单向通知，所以这种方法只适用于非常简单的逻辑且并发量不太大的场景，一旦逻辑稍微复杂一点，这种方法就有点捉襟见肘。
*/
