/*
224. 基本计算器
2022-12-25
link: https://leetcode.cn/problems/basic-calculator/
question:
	给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
	注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
answer:
	栈实现计算。
	栈只存数字.
	粗暴想法，仅仅通过25/44:
	- 空格影响可以直接全部去除，难点在括号。
	- 如果括号所有去除，会影响‘-’的一元运算有效性！

	参考：括号展开+栈
	一对括号是一层，一层层计算。外层要等内层解决了才能发挥作用。
	https://leetcode.cn/problems/basic-calculator/solution/si-lu-jian-dan-yi-ci-bian-li-dai-ma-jian-rzjn/
*/
package main

import "unicode"

func calculate(s string) int {
	stack := make([]int, len(s)) // 先申请最大内存，避免后面扩容等操作。
	sign := 1                    // 1为正，-1为负
	res := 0                     // res为当前层的计算结果
	num := 0                     // num保存一个一位或多位的数字
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			// 构建多位数
			j := i
			num = 0
			for j < len(s) && unicode.IsDigit(rune(s[i])) {
				num = num*10 + int(s[j]-'0')
				j++
			}
			res += sign * num // 累加得到最近的结果
			i = j - 1         // i要跟上j,由于for循环最后j多加一次指向的不是数字，而外层for循环i还会再加一次。
		case '(':
			stack = append(stack, res)  //将最新结果入栈
			stack = append(stack, sign) // 将最近的符号入栈
			res = 0                     // 重置res,因为要计算新的层级
			sign = 1                    // 重置sign
		case ')':
			sign = stack[len(stack)-1]    // 获取栈顶的正负符号
			preRes := stack[len(stack)-2] // 获取上一层的计算结果
			stack = stack[:len(stack)-2]  // 让栈顶的两个元素出站
			res = preRes + sign*res       //更新最新结果
		case '+':
			sign = 1
		case '-':
			sign = -1
		}
	}
	return res
}

// 粗暴想法，不行
// func calculate(s string) int {
// 	// 去除所有空格
// 	s = strings.Replace(s, " ", "", -1) // 通过替换，去除字符串中的指定字符串
// 	// 去除左右括号
// 	s = strings.Replace(s, "(", "", -1)
// 	s = strings.Replace(s, ")", "", -1)
// 	// 栈实现计算
// 	var stack []int = []int{}
// 	var preSign byte = '+' // 关键,数字前的符号标志
// 	var num int = 0        // 当前数字
// 	for i := 0; i < len(s); i++ {
// 		isDigital := s[i] >= '0' && s[i] <= '9'
// 		if isDigital {
// 			num = num*10 + int(s[i]-'0')
// 		}
// 		if !isDigital || i == len(s)-1 {
// 			// 当前不是数字时，就只能是符号，所以要根据保存的符号和上一个数字添加到栈里。然后在记录当前符号
// 			switch preSign {
// 			case '+':
// 				stack = append(stack, num)
// 			case '-':
// 				stack = append(stack, -num)
// 			}
// 			// 下一轮开始
// 			// if s[i] == '(' || s[i] == ')' {
// 			// 	continue
// 			// }
// 			preSign = s[i]
// 			num = 0
// 		}
// 	}
// 	// 统计结果
// 	var res int = 0
// 	for i := 0; i < len(stack); i++ {
// 		res += stack[i]
// 	}
// 	return res
// }
