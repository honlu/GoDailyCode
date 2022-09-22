/*
2022-9-22
question: 在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

answer:将一个字符串以中间字符为轴，前后翻转，即将str[len]赋值给str[0],将str[0] 赋值 str[len]。
*/
package main

import (
	"fmt"
	"log"
)

func main() {
	var s string // 注意string是不能更改的
	fmt.Scan(&s)
	t := []rune(s) // string转为[]rune, 不能写成s = []rune(s)
	l := len(t)
	if l > 5 { // 这里为了测试，5000改为5
		log.Panic("strings 超过5") // panic
	}
	for i := 0; i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	s = string(t)
	fmt.Println(s)
}
