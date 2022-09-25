/*
2022-9-25
question:
	两个goroutine使用，一个输出数字，一个输出字母
answer:
	进一步学习channel,go,sync.WaitGroup
*/
package main

import (
	"fmt"
	"sync"
)

var (
	wg    = sync.WaitGroup{}
	chNum = make(chan bool)
	chAlp = make(chan bool)
)

func main() {
	go func() {
		i := 1
		for {
			<-chNum
			fmt.Printf("%d%d", i, i+1)
			i = i + 2
			chAlp <- true
		}
	}()
	wg.Add(1)
	go func() {
		i := 'a'
		for {
			<-chAlp
			fmt.Print(string(i), string(byte(i)+1))
			i = i + 2
			if i > 'z' {
				wg.Done()
				// return  // 如果没有return，打印数组可能会再执行一次.或者加else.就可以多避免打印一次数字
			} else {
				chNum <- true
			}
		}
	}()
	chNum <- true
	wg.Wait()
}
