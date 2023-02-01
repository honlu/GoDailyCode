package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch // 接收到数据后，阻塞当前协程等待发送方发送数据。（本main函数结束后，当前goroutine也结束）
	}()
	ch <- true // 无缓冲，发送方阻塞知道接收方接受到数据，然后执行下面的语句
	fmt.Println("时间花费：", time.Now().Sub(t).Seconds())
	time.Sleep(time.Second * 2)
}
