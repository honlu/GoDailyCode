/*
10. 正则表达式匹配
2022-11-6
2023-4-26 updated, labeled as star by lu
link:https://leetcode.cn/problems/regular-expression-matching
question:
	给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
	所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
answer:
	字符串匹配，经典DP问题。动态规划五步曲：
	1、dp数组和下标含义：
		定义动态数组dp[m+1][n+1](n为p的长度，m为字符串的长度)
		dp[i][j]表示s的前i个字符和p的前j个字符能否匹配。
	2、递推公式
		在进行推导时，s 中的字符是固定不变的，我们考虑 p 的第 j 个字符与 s 的匹配情况：
		1、p[j]是一个小写字母a-z,则s[i]必须也为同样的小写字母方能完成匹配，
			即if s[i]== p[j]: dp[i][j]=dp[i-1][j-1] else dp[i][j] = false
		2、p[j]='.'时，则默认s[i]=p[j]
		3、p[j]='*'时，还要区分两种情况：
			若p[j-1]=s[i] ,则dp[i][j]=dp[i-1][j]
			若p[j-1]='.', 则dp[i][j]= dp[i][j-2] 注意：
			若p[j-1]!=s[i],则dp[i][j] = dp[i][j-2] 注意：
	3、初始化
		dp[0][0]=true,最终的答案为dp[m][n]
	4、遍历顺序
		从上到下，从左到右
	5、举例推导
*/
func isMatch(s string, p string) bool {
	// 创建dp数组和初始化
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	// 遍历和更新
	for i := 0; i <= m; i++ { // s字符串能为空串，所以从0开始
		for j := 1; j <= n; j++ { // p字符串不能为空串，所以从1开始
			// 这里简化了，不太好理解
			if p[j-1] == '*' {
				if i != 0 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else {
				if i != 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}
	return dp[m][n]
}