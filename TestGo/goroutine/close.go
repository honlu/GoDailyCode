/*
如何关闭goroutine?
- 关闭channel
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func test1() {
	ch := make(chan string, 6)
	go func() {
		for {
			v, ok := <-ch //  这种接收channel的方式，可以利用其第二个参数进行判别，当关闭 channel 时，就根据其返回结果跳出。
			if !ok {
				fmt.Println("结束")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- "煎鱼还没进锅里..."
	ch <- "煎鱼进脑子里了！"
	close(ch)
	time.Sleep(time.Second)
}

func test2() {
	ch := make(chan string, 6)
	done := make(chan struct{})
	// 启动两个协程
	go func() {
		for {
			select {
			case ch <- "脑子进煎鱼了":
			case <-done:
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond) // 100ms
		}
	}()

	go func() {
		time.Sleep(1 * time.Second) // 1s后通过channel通知另一个协程关闭另一个channel以及退出
		done <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("接收到的值: ", i)
	}
	fmt.Println("结束")
}

func test3() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background()) // 借助context关闭goroutine

	go func(ctx context.Context) {
		// go func() { // 上一个协程退出，这个协程不退出
		// 	for {
		// 		fmt.Println("测试子协程")
		// 		time.Sleep(300 * time.Millisecond)
		// 	}
		// }()
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return // 这里return加不加都一样，不过为了避免错误加上挺好
			default:
				fmt.Println("煎鱼还没到锅里...")
			}

			time.Sleep(500 * time.Millisecond)
		}

	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("结束")
	time.Sleep(1 * time.Second)
}

func main() {
	// test1()
	// test2()
	test3()
}
