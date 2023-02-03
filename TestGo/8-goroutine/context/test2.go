/*
context.WithValue如果需要往子协程中传递参数，可以使用 context.WithValue()。
*/
package main

import (
	"context"
	"fmt"
	"time"
)

type Options struct{ Interval time.Duration }

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("options").(*Options) // 在子协程中，使用 ctx.Value("options") 获取到传递的值，读取/修改该值
			time.Sleep(op.Interval * time.Second) // 等待x秒
		}
	}
}

func test1() { // WithCancel
	ctx, cancel := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx, "options", &Options{1}) // context.WithValue() 创建了一个基于 ctx 的子 Context，并携带了值 options。
	// 多个协程
	go reqTask(vCtx, "worker1")
	go reqTask(vCtx, "worker2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

/*
如果需要控制子协程的执行时间，可以使用 context.WithTimeout 创建具有超时通知机制的 Context 对象。
*/
func reqTask2(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func test2() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // WithTimeout()的使用与 WithCancel() 类似，多了一个参数，用于设置超时时间。
	go reqTask2(ctx, "worker1")
	go reqTask2(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel") // 因为超时时间设置为 2s，但是 main 函数中，3s 后才会调用 cancel()，因此，在调用 cancel() 函数前，子协程因为超时可能已经退出了。
	cancel()
	time.Sleep(3 * time.Second)
}

/*
超时退出可以控制子协程的最长执行时间，那 context.WithDeadline() 则可以控制子协程的最迟退出时间。
*/
func reqTask3(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name, ctx.Err())
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func test3() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	go reqTask3(ctx, "worker1")
	go reqTask3(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(3 * time.Second)
}

func main() {
	// test2()
	test3()
}
