/*
通过channel和sync.WaitGroup方式限制协程数量.
注意，必须使用make创建channel。
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type Glimit struct { // 这里定义groutine的限制结构体
	n int
	c chan struct{}
}

// initialization Glimit struct 初始化结构体
func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

// Run f in a new goroutine but with limit.
// 结构体的方法
func (g *Glimit) Run(f func()) {
	g.c <- struct{}{} // 在创建一个goroutine前，读入一个channel，来计数。
	go func() {
		f()
		<-g.c // 执行完一个goroutinne时，接收一个channel.相当于一个goroutine结束，channel减少一个
	}()
}

var wg = sync.WaitGroup{}

func main() {
	number := 10
	g := New(2) // 限制最多两个goroutine
	for i := 0; i < number; i++ {
		wg.Add(1) // 在新建一个goroutine前，计数器+1
		value := i
		goFunc := func() {
			// 做一些业务逻辑处理
			fmt.Printf("go func: %d\n", value)
			time.Sleep(time.Second)
			wg.Done()
		}
		g.Run(goFunc)
	}
	wg.Wait()
}
