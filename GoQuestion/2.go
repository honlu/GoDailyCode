/*
2022-9-22
question:确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

answer:使用的是golang内置方法`strings.Count`,可以用来判断
在一个字符串中包含的另外一个字符串的数量

这里有几个重点，第一个是`ASCII字符`，`ASCII字符`字符一共有256个，其中128个是常用字符，可以在键盘上输入。128之后的是键盘上无法找到的。

然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使用额外的储存结构，且字符串小于等于3000。

如果允许其他额外储存结构，这个题目很好做。如果不允许的话，可以使用golang内置的方式实现。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	var res bool
	res = true
	// 判断字符串是否全都不同
	if len(s) > 3000 { // 字符串长度不能超过3000，len(s) 等效于 strings.Count(s,"")
		res = false
	}
	for _, v := range s { // range string注意v是rune类型
		if v > 127 { // 判断是否是ascii字符
			res = false
		}
		if strings.Count(s, string(v)) > 1 { // 字符是否超过1
			res = false
		}
	}
	fmt.Println(res)
}
