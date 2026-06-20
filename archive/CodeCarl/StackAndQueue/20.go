/*
20. 有效的括号
2023-1-12
link: https://leetcode.cn/problems/valid-parentheses/
question:
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，
	判断字符串是否有效。
	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		每个右括号都有一个对应的相同类型的左括号。

answer:
	(([)]) 这个例子是否true呢？是false
	利用栈匹配：栈中存左括号。
	哈希匹配
*/
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	// 栈声明
	stack := make([]byte, 0)
	// 哈希
	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	// 遍历
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}