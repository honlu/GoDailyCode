/*
context.WithCancel
context.WithCancel() 创建可取消的 Context 对象，即可以主动通知子协程退出。

使用 Context 改写，效果与 select+chan 相同。

例子：控制单个协程
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // 在子协程中，使用 select 调用 <-ctx.Done() 判断是否需要退出。
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func contral2() {
	ctx, cancel := context.WithCancel(context.Background())
	// 产生多个协程
	go reqTask(ctx, "worker1")
	go reqTask(ctx, "worker2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)

}

func contral1() {
	ctx, cancel := context.WithCancel(context.Background()) // context.Backgroud() 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context。
	// context.WithCancel(parent) 创建可取消的子 Context，同时返回函数 cancel。
	go reqTask(ctx, "worker1")
	time.Sleep(3 * time.Second)
	cancel() // 主协程中，调用 cancel() 函数通知子协程退出。
	time.Sleep(3 * time.Second)
}

func main() {
	// contral1()  // context控制一个协程

	contral2() // context控制多个协程

	// context.WithValue 如果需要往子协程中传递参数，可以使用 context.WithValue()。
}
