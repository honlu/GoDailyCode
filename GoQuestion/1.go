/*
2022-9-16 by honlu
question： 实现两个goroutine交替打印数字或字母
涉及类型或函数等：channel、sync.WaitGroup、go func(){}、select
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan bool), make(chan bool) // 两个channel，分别作为控制数字和字母的通信
	wait := sync.WaitGroup{}

	go func() { // 打印数字的goroutine
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true // 通知字母打印
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true // 通知数字打印
			}
		}
	}(&wait)
	number <- true // 通知数字打印
	wait.Wait()    // 阻塞主goroutine，直到子goroutine结束
}
