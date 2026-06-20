/*
17. 最长回文子序列
2022-11-2
link: https://leetcode.cn/problems/longest-palindromic-subsequence/
question:
	给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。
	可以假设 s 的最大长度为 1000 。
answer:
	回文子串是要连续的，回文子序列可不是连续的！
	回文子串，回文子序列都是动态规划经典题目。
	动态规划五步曲：
	1、dp数组和下标含义：
		dp[i][j]：字符串s在[i, j]范围内最长的回文子序列的长度为dp[i][j]。
	2、递推公式
		如果s[i]与s[j]相同，那么dp[i][j] = dp[i + 1][j - 1] + 2;
		如果s[i]与s[j]不相同，说明s[i]和s[j]的同时加入 并不能增加[i,j]区间回文子串的长度
			dp[i][j] = max(dp[i + 1][j], dp[i][j - 1]);
	3、初始化
		当i与j相同，那么dp[i][j]一定是等于1的，即：一个字符的回文子序列长度就是1。
		其他情况dp[i][j]初始为0就行，这样递推公式：dp[i][j] = max(dp[i + 1][j], dp[i][j - 1]); 中dp[i][j]才不会被初始值覆盖。
	4、遍历顺序
		遍历i的时候一定要从下到上遍历，这样才能保证，下一行的数据是经过计算的。
	5、举例推导
*/

func longestPalindromeSubseq(s string) int {
	// dp数组和初始化
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	// 遍历
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}