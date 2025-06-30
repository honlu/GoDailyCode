package stack

/*
69-有效的括号
https://leetcode.cn/problems/valid-parentheses/description/?envType=study-plan-v2&envId=top-100-liked
*/

func isValid(s string) bool {
	// 典型的栈，左符号进展，右符号在顶部遇到合适的左符合，栈元素出
	stack := []byte{}
	flag := map[byte]byte{
		']': '[',
		')': '(',
		'}': '{',
	}
	if len(s)%2 == 1 { // 注意是否成对
		return false
	}
	for _, item := range []byte(s) {
		if item == '(' || item == '{' || item == '[' {
			stack = append(stack, item)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != flag[item] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) > 0 { // 注意栈是否元素出完
		return false
	}
	return true
}
