/*
使用channel进行通信,并发控制goroutine.
实现优雅等待所有子goroutine完全结束之后主进程才结束退出.
借助了标准库sync里的Waitgroup，这是一种控制并发的方式，可以实现对多goroutine的等待.
如果并发启动了多个子协程，需要等待所有的子协程完成任务，WaitGroup 非常适合于这类场景.

然而WaitGroup 只是傻傻地等待子协程结束，但是并不能主动通知子协程退出。
假如开启了一个定时轮询的子协程，有没有什么办法，通知该子协程退出呢？这种场景下，可以使用 select+chan 的机制。
子协程使用 for 循环定时轮询，如果 stop 信道有值，则退出，否则继续轮询。
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func consumer(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("exit sub goroutine")
			return
		default:
			fmt.Println("running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	stop := make(chan bool) // 可以通知子协程提前结束。
	var wg sync.WaitGroup   // 定义一个全局WaitGroup
	// Spawn example consumers
	for i := 0; i < 3; i++ {
		wg.Add(1) // 每当启动一个goroutine前，计数器+1
		go func(stop <-chan bool) {
			defer wg.Done() // 保证goroutine结束，计数器-1
			consumer(stop)
		}(stop)
	}
	waitForSignal()
	close(stop) // 通知全部子协程结束运行。
	fmt.Println("stopping all jobs!")
	wg.Wait() // 主线程等待所有goroutine结束
}
func waitForSignal() { // 这个干什么的呢？
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
}

/*
sync.WaitGroup:它的源码里实现了一个类似计数器的结构，记录每一个在它那里注册过的协程，
然后每一个协程完成任务之后需要到它那里注销，然后在主进程那里可以等待直至所有协程完成任务退出。

使用步骤：
1. 创建一个Waitgroup的实例wg；
2. 在每个goroutine启动的时候，调用wg.Add(1)注册；
3. 在每个goroutine完成任务后退出之前，调用wg.Done()注销。
4. 在等待所有goroutine的地方调用wg.Wait()阻塞进程，知道所有goroutine都完成任务调用wg.Done()注销之后，Wait()方法会返回。

该示例程序是一种golang的select+channel的典型用法。
继续分析select：
	golang 的 select 机制可以理解为是在语言层面实现了和 select, poll, epoll 相似的功能：
	监听多个描述符的读/写等事件，一旦某个描述符就绪（一般是读或者写事件发生了），就能够将发生的事件通知给关心的应用程序去处理该事件。
	golang 的 select 机制是，监听多个channel，每一个 case 是一个事件，可以是读事件也可以是写事件，随机选择一个执行，可以设置default,
	它的作用是：当监听的多个事件都阻塞住会执行default的逻辑。
*/
