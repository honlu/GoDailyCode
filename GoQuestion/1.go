/*
2022-9-16 by honlu
2023-1-3 update by zll
question：
	实现两个goroutine交替打印数字或字母.
	使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母。
	最终效果如下：
	12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
answer:
	涉及类型或函数等：channel、sync.WaitGroup、go func(){}、select

*/
package main

import (
	"fmt"
	"sync"
)

/*
解题思路1：只是用channel，不适用sync.WaitGroup，goroutine的结束控制通过channel来控制
*/
func solve1() {
	number := make(chan bool) // 控制数字
	letter := make(chan bool) // 控制字母
	done := make(chan bool)   // 终止信号chan

	// 打印数字
	go func() {
		i := 1
		for {
			select {
			case <-number: // 收到控制数字的信号，就输出
				fmt.Printf("%d%d", i, i+1) // 如果fmt.Print(i, i+1)直接打印两个数字，会有空格
				i += 2
				// 接着打印字母
				letter <- true
			}
		}
	}()
	// 打印字母
	go func() {
		j := 'A'
		for {
			select {
			case <-letter:
				if j > 'Z' { // 当字母打印结束，回到主goroutine
					done <- true
				} else {
					fmt.Print(string(j), string(j+1))
					j += 2
					number <- true // 交替打印，去打印数字
				}
			}
		}
	}()
	// 开始打印数字
	number <- true

	// 阻塞主goroutine，等待全部goroutine结束
	for {
		select {
		case <-done:
			return
		}
	}
}

/*
解题思路2：使用sync.WaitGroup控制goroutine的退出
*/
func solve2() {
	letter, number := make(chan bool), make(chan bool) // 两个channel，分别作为控制数字和字母的通信
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() { // 打印数字的goroutine
		i := 1
		for {
			select {
			case <-number:
				fmt.Printf("%d%d", i, i+1)
				i += 2
				if i == 29 {
					wg.Done()
				}
				letter <- true // 通知打印字母的goroutine打印
			}
		}
	}()
	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wg.Done()
					return
				} else {
					fmt.Print(string(i), string(i+1))
					i += 2
					number <- true // 通知打印数字的goroutine打印
				}
			}
		}
	}()

	number <- true // 通知数字打印
	wg.Wait()      // 阻塞主goroutine，直到全部goroutine结束
}

func main() {
	// solve1()
	solve2()
}
