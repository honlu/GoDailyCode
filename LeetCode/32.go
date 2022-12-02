/*
32. 最长有效括号
link: https://leetcode.cn/problems/longest-valid-parentheses/
question:
	给你一个只包含 '(' 和 ')' 的字符串，
	找出最长有效（格式正确且连续）括号子串的长度。
answer:
	栈实现！
*/
// 参考题解，栈中存右括号的下标！
func longestValidParentheses(s string) int {
	var stack []int
	max := 0
	stack = append(stack, -1) // 从下标-1开发
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1] // 这里很牛！遇到左括号就出栈
			if len(stack) == 0 {         // 则下标都没有，新的一轮从i开发
				stack = append(stack, i)
			} else {
				if i-stack[len(stack)-1] > max { // 注意这里i减去栈顶部元素的下标！这样就可以避免右括号过多导致的问题
					max = i - stack[len(stack)-1]
				}
			}
		}
	}
	return max
}

// 一下代码有问题，栈中存右括号的话，只能通过156/231用例！bug出现在右括号过多的情况！处理有些麻烦。
// func longestValidParentheses(s string) int {
// 	var stack []byte
// 	max := 0
// 	flag := 0
// 	for i := 0; i < len(s); i++ {
// 		if len(stack) == 0 {
// 			if s[i] == '(' {
// 				stack = append(stack, s[i])
// 			} else {
// 				if flag > max {
// 					max = flag
// 				}
// 				flag = 0
// 			}
// 		} else {
// 			if s[i] == '(' {
// 				stack = append(stack, s[i])
// 			} else {
// 				flag += 2
// 				stack = stack[:len(stack)-1]
// 				if flag > max {
// 					max = flag
// 				}
// 			}
// 		}

// 	}
// 	return max
// }
