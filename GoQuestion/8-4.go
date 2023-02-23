/*
以下代码:
编译失败，值类型 Student{} 未实现接口People的方法，不能定义为 People类型。
在 golang 语言中，Student 和 *Student 是两种类型，第一个是表示 Student 本身，第二个是指向 Student 的指针。
*/
package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Student{} // 这里值类型 Student{} 未实现接口People的方法，不能定义为 People类型。
	// 因为是*Student类型实现了接口的方法.所以改成&Student{}就可以了
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
