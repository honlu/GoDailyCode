/*
2022-9-22
question:
	给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
	这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。
	保证两串的长度都小于等于5000。
answer:
	首先要保证字符串长度小于5000。
		不需要排序，只需要确定两个字符串中字符个数都一样多就可以了
	之后只需要一次循环遍历s1中的字符在s2是否都存在即可。
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	l1, l2 := len(s1), len(s2)
	if l1 > 5000 || l2 > 5000 || l1 != l2 { // 首先长度判断，不超过5000，两个字符串长度要一样
		fmt.Println(false)
		return
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) { // 两个字符串中每个字符个数要一样
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
}
