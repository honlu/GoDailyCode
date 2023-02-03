/*
select+chan 的机制
	开启了一个定时轮询的子协程，通知该子协程退出
*/
package main

import (
	"fmt"
	"time"
)

var stop chan bool

func reqTask(name string) {
	for {
		select {
		case <-stop:
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop = make(chan bool)
	go reqTask("worker1")
	time.Sleep(3 * time.Second) // 定时3秒后，关闭子协程
	stop <- true
	time.Sleep(3 * time.Second) // 定时3秒后，主线程结束
}
