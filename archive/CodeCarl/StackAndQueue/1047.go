/*
1047. 删除字符串中的所有相邻重复项
2023-1-13
link: https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
question:
	给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
	在 S 上反复执行重复项删除操作，直到无法继续删除。
	在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

answer:
	栈思想，当栈顶元素和要入栈元素相等时，栈顶和入栈元素同时删除。
*/
func removeDuplicates(s string) string {
	// 栈
	stack := make([]byte, 0)
	// 遍历
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] != stack[len(stack)-1] { // 不等，入栈
			stack = append(stack, s[i])
		} else { // 相等，出栈
			stack = stack[:len(stack)-1]
		}
	}
	// 返回栈中剩余元素
	return string(stack)
}