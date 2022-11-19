/*
415. 字符串相加
2022-11-19
link: https://leetcode.cn/problems/add-strings/
question:
	给定两个字符串形式的非负整数 num1 和num2 ，
	计算它们的和并同样以字符串形式返回。
answer:
	注意：不能适用内建的用于处理大整数的库，也不可以字符串转为整数！
	方法：模拟遍历进行添加
*/
package main

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	// num1一定要是长度小的值
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	var res string
	var flag int // 注意：flag:= '0' 则flag是rune类型,不是byte类型
	var temp int
	i := 1
	for ; i <= len(num1); i++ {
		temp = int(num1[len(num1)-i]-'0') + int(num2[len(num2)-i]-'0') + flag // 注意byte类型，不能进行加操作！
		if temp > 9 {
			temp = temp % 10
			flag = 1
		} else {
			flag = 0
		}
		res += strconv.Itoa(temp) // 整数转字符串
	}
	for ; i <= len(num2); i++ {
		temp := int(num2[len(num2)-i]-'0') + flag
		if temp > 9 {
			temp = temp % 10
			flag = 1
		} else {
			flag = 0
		}
		res += strconv.Itoa(temp)
	}
	if flag != 0 {
		res += strconv.Itoa(flag)
	}
	// 字符串反转
	res = reverse(res)
	return res
}

func reverse(temp string) string {
	s := []byte(temp)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	s1 := "11"
	s2 := "123"
	addStrings(s1, s2)
}
