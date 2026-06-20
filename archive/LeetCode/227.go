/*
227. 基本计算器2
2022-12-11
link: https://leetcode.cn/problems/basic-calculator-ii/
question:
	给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
	整数除法仅保留整数部分。

	你可以假设给定的表达式总是有效的。
	所有中间结果将在 [-231, 231 - 1] 的范围内。

	注意：不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
answer:
	注意整数除法仅保留整数部分，所以不用考虑小数等类型转换复杂操作。
	利用栈思想模拟操作:
		栈中只保存数值，加号就直接将数字压入栈；减号就将数字的相反数压入栈；
		乘除就计算好数字与栈顶元素运算结果，并替换栈顶元素。
	注意数字可能不是个位数！
*/
package main

import "strconv"

func calculate(s string) int {
	var stack []int        // 栈
	num := 0               // 当前数字
	var preSign byte = '+' // 数字前的符号标志
	var res int
	if len(s) == 1 {
		res, _ = strconv.Atoi(s)
		return res
	}
	for i := 0; i < len(s); i++ {
		isDigital := s[i] >= '0' && s[i] <= '9'
		if isDigital {
			num = num*10 + int(s[i]-'0') // num*10是为了保证连续几个都是数字
		}
		if !isDigital && s[i] != ' ' || i == len(s)-1 { // 当前s[i]不是数字时，即符号,要开始计算上一个操作运算了。
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			// 下一轮操作开始
			preSign = s[i]
			num = 0
		}
	}
	// 计算栈中剩余数字相加
	for _, val := range stack {
		res += val
	}
	return res
}
