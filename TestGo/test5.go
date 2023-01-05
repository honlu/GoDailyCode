/*
Golang对全局变量加锁同步解决资源访问共享问题.
map的并发访问.
使用Go协程来同时并发计算多个数字（1-200）的（0-i）累加，然后存储在数组当中
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//声明一个全局互斥锁
	lock sync.Mutex
)

func test(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i //这里我将阶乘改成求和，防止数据溢出
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10) // 为了主线程等待所有goroutine结束

	for i, v := range myMap { // map访问是随机的
		fmt.Printf("myMap[%d]=%d\n", i, v)
	}
}

/*
上述解决方案不太完美，有其缺陷：

（1）主线程在等待所有goroutine全部完成的时间很难确定

（2）如果主线程休眠时间过长，就会加长等待时间，如果等待时间短了，还可能会有goroutine因为主线程的退出而被销毁

（3）通过全局变量加锁同步来实现通讯，也不利于多个协程对全局变量的读写操作
*/
