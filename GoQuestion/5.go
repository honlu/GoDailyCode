/*
2022-9-22
question:
	将字符串中的空格全部替换为“%20”。
	假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于1000)，
	同时保证字符串由【大小写的英文字母组成，注意这是提示范围了！】。
	给定一个string为原始的串，返回替换后的string。
answer:
	使用了golang内置方法`unicode.IsLetter`判断字符是否是字母
	之后使用`strings.Replace`来替换空格
关于字符串修改：
	注意：字符串是无法被修改的，只能复制原字符串，在复制的版本上修改
	方法1：转换为[]byte()
	方法2：转换为[]rune()
	方法3：新字符串代替原字符串的子字符串,用strings包中的strings.Replace()
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	var s string
	// fmt.Scan(&s) // 注意这里就不能使用Scan输入。
	// 需要获取整行内容，可能包含空格或其他特殊符号等，使用bufio.NewReader(os.Stdin) ReadLine()
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine() // 返回的line是[]byte类型
	if err != nil {
		log.Panic(err.Error())
	}
	s = string(line)
	// fmt.Println(s)  // 测试
	if len(s) > 1000 {
		log.Panic("超出长度")
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			log.Panic("字符只能是空格或是大小写英文字母")
		}
	}
	s = strings.Replace(s, " ", "%20", -1) // 替换字符串中空格
	fmt.Println(s)
}
