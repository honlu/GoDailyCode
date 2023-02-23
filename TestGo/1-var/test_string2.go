// 测试string不可更改的原因
// s[i]某个字符是只读的、不可更改。s整体是可以更改的，即指向新的字符地址。
package main

import "fmt"

func main() {
	var s string = "hello"
	fmt.Println(s)
	fmt.Println(&s)

	s = "welcome"
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(&s)
}
