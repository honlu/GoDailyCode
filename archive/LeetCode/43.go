/*
43. 字符串相乘
2022-12-1
link: https://leetcode.cn/problems/multiply-strings/
question:
	给定两个以字符串形式表示的非负整数 num1 和 num2，
	返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
answer:
	模拟操作字符串，迭代实现
*/
package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	s2 := len(num2) - 1
	res := ""
	if num1 == "0" || num2 == "0" { // 注意提示中特殊情况：num1和num2都不包含任何前导零，除了数字0本身。
		return "0" // 若无这个判断，后面的结果可能会出现"0000",和想要的"0"不一致
	}
	for s2 >= 0 {
		flag := 0
		temp := ""
		// 补零
		for i := 0; i < len(num2)-1-s2; i++ {
			temp += "0"
		}
		// 字符串和s2位置的字母相乘
		for i := len(num1) - 1; i >= 0 || flag != 0; i-- {
			var y int
			x := int(num2[s2] - '0')
			if i >= 0 {
				y = int(num1[i] - '0')
			}
			result := x*y + flag
			temp = strconv.Itoa(result%10) + temp
			flag = result / 10
		}
		fmt.Printf("res:%s,temp:%s\n", res, temp)
		// 上一次结果res和本次结果temp相加[注意相加操作]
		add := 0
		ans := "" // 保存结果
		for i, j := len(res)-1, len(temp)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
			var x, y int
			if i >= 0 {
				x = int(res[i] - '0')
			}
			if j >= 0 {
				y = int(temp[j] - '0')
			}
			result := x + y + add
			ans = strconv.Itoa(result%10) + ans
			add = result / 10
		}
		res = ans // 更新res
		s2--
	}
	return res
}

func main() {
	// s1, s2 := "123", "456"
	s1, s2 := "123789", "456"
	res := multiply(s1, s2)
	fmt.Println(res)
}
