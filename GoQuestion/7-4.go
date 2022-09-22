/*
2022-9-22
question:
	自定义函数和已有内置函数，循环调用问题
answer:
	循环调用会导致fatal error: stack overflow
	在golang中String() string 方法实际上是实现了String的接口的，该接口定义在fmt/print.go 中：

		type Stringer interface {
			String() string
		}

	在使用 fmt 包中的打印方法时，如果类型实现了这个接口，
	会直接调用。
	而题目中打印 p 的时候会直接调用 p 实现的String() 方法，
	然后就产生了循环调用。(还有些迷糊！)
*/
package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	p.String()
}
