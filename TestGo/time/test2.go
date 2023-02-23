/*
每秒执行一次proc并保证程序不退出
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second * 1) //定时器
	for {
		select {
		case <-t.C:
			func() {
				defer func() { // defer一定要在proc执行之前
					if err := recover(); err != nil {
						fmt.Println("recover", err)
					}
				}()
				proc()
			}()
		}
	}

}

func proc() {
	panic("ok")
}
