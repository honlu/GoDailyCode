/*
151. 反转字符串中的单词
2023-1-11
link: https://leetcode.cn/problems/reverse-words-in-a-string/
question:
	给你一个字符串 s ，请你反转字符串中 单词 的顺序。
	单词 是由非空格字符组成的字符串。
	s 中使用至少一个空格将字符串中的 单词 分隔开。
	返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

answer:
	注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。
	返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

	暴力模拟。
	还可以优化：不要使用辅助空间，空间复杂度要求为O(1)。
*/
// 暴力模拟
func reverseWords(s string) string {
	// 首先将全部单词保存到一个字符串数组中
	temp := ""
	st := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			temp += string(s[i])
		} else if s[i] == ' ' && temp != "" {
			st = append(st, temp)
			temp = ""
		}
	}
	if temp != "" {
		st = append(st, temp)
	}
	// 然后按照要求倒着拼接字符串数组
	res := st[len(st)-1]
	for i := len(st) - 2; i >= 0; i-- {
		res += " " + st[i]
	}
	return res
}