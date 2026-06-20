package stack

import "strings"

/*
71-394.字符串解码
https://leetcode.cn/problems/decode-string/description/?envType=study-plan-v2&envId=top-100-liked
*/
/*
思路：双栈解法
*/
func decodeString(s string) string {
	// 从前往后遍历
	// 栈: 双栈使用数字栈，字符串栈。
	num := 0              // 刚读到的重复次数
	result := ""          // 当前字符串片段
	var numStack []int    // 入栈每一层的重复次数
	var strStack []string // 入栈每一层进入“[”的“已构建前缀”
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '[' { // 入栈，把当前层的 result 和 num 入栈，然后把当前层清空，开始构建下一层内容
			strStack = append(strStack, result) // 注意
			result = ""
			numStack = append(numStack, num)
			num = 0
		} else if s[i] == ']' { // 出栈，说明当前层片段结束。
			// 弹出外层的前缀 str 和次数 count，把“当前层 result 重复 count 次”，接回外层前缀上，得到新的 result（相当于退回上一层）。
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1] // 出栈
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			result = str + strings.Repeat(result, count)
		} else { // 添加字符串
			result += string(s[i])
		}
	}
	return result
}
