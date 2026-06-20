/*
459. 重复的子字符串
2023-1-11
link: https://leetcode.cn/problems/repeated-substring-pattern/
question:
	给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
answer:
	1. 暴力解法：一个for循环获取 子串的终止位置，然后判断子串是否能重复构成字符串，
		又嵌套一个for循环，所以是O(n^2)的时间复杂度。
	还可以优化：
	2. 移动匹配
	3. KMP
*/
// 暴力解法：假设子串长度为i,则s[j]=s[j-i]
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	var res bool
	for i := 1; i*2 <= n; i++ { // 子串长度不能超过n/2
		if n%i == 0 { // 原长度一定是子串长度的整数倍
			res = true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] { // 当前假设子串长度i不对，break换下一个
					res = false
					break
				}
			}
			if res { // 找到一个就可以提前结束了
				return res
			}
		}
	}
	return res
}