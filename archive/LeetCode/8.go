package leetcode

import (
	"math"
	"strings"
)

/*
8. 字符串转换整数（atoi）
2022-11-26
link:https://leetcode.cn/problems/string-to-integer-atoi/
question:
	请你来实现一个 myAtoi(string s) 函数，
	使其能将字符串转换成一个 32 位有符号整数
	（类似 C/C++ 中的 atoi 函数）。
answer:
	字符串处理，函数调用.
	注意：有个特殊测试用例“9223372036854775808”,在myAtoi计算res结果为负的。
	这是因为：int36的整数最大值是9223372036854775807，所以在计算机中会导致负号！
// */
// package main

// import (
// 	"fmt"
// 	"math"
// 	"strings"
// 	"unicode"
// )

// func myAtoi(s string) int {
// 	var res int
// 	// 1. 去除前导空格
// 	s = strings.Trim(s, " ")
// 	if len(s) == 0 { // 注意特殊情况
// 		return 0
// 	}
// 	// 2. 保留正负号标志
// 	var flag int = 1
// 	if s[0] == '-' { // 测试用例里会有+号吗?有可能
// 		flag = -1
// 		s = s[1:]
// 	} else if s[0] == '+' {
// 		s = s[1:]
// 	}
// 	// 3. 读取数字： 数字用什么保存呢？ 4. 抹掉数字前置0的部分 5. 判断数组是否超出32位范围
// 	for _, val := range s {
// 		if unicode.IsNumber(val) {
// 			res = 10*res + int(val-'0') // 这一步精妙！
// 			if flag*res < math.MinInt32 {
// 				return math.MinInt32
// 			} else if flag*res > math.MaxInt32 {
// 				return math.MaxInt32
// 			}
// 		} else {
// 			break
// 		}
// 	}
// 	// 6. 返回结果
// 	return res * flag
// }

// func main() {
// 	s := "18446744073709551617"
// 	res := myAtoi(s)
// 	fmt.Println(res)
// }

/*
模拟题
*/
func myAtoi(s string) int {
	// 模拟题，重点「边界问题」
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}
	var res int
	var flag int = 1
	if s[0] == '-' {
		flag = -1
	}
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	for i := 0; i < len(s); i++ {
		if s[i] > '9' || s[i] < '0' {
			break
		}
		res = res*10 + int(s[i]-'0')
		if flag*res > math.MaxInt32 { // 重点在这里flag*res
			return math.MaxInt32
		} else if flag*res < math.MinInt32 {
			return math.MinInt32
		}
	}

	return res * flag
}
