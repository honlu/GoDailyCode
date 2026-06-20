/*
28. 找出字符串中第一个匹配项的下标
2023-1-11
link: https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
question:
	给你两个字符串 haystack 和 needle ，
	请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标
	（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。

answer：
	KMP 经典题目!
	1. 可以暴力遍历解决
	2. 构建前缀表，KMP字符串匹配。（难点）
*/
// 暴力遍历
func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	for i := 0; i <= m-n; i++ { // 注意这里等于号
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1

}