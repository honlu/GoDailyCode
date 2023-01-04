/*
输出: BBBBBBB
跟上一题一样，不同的是*Student 的定义后本身没有初始化值，所以 *Student 是 nil的，但是*Student 实现了 People 接口，接口不为 nil。
这里就涉及一些反射的知识了,一个变量不仅仅涉及他自身的类型,还有值
*/

package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
