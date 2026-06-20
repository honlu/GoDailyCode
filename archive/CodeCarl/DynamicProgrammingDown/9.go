/*
9. 最长公共子序列
2022-11-2
link: https://leetcode.cn/problems/longest-common-subsequence/
question:
	给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
	一个字符串的 子序列 是指这样一个新的字符串：
	它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
answer:
	注意：在于这里不要求是连续的了，但要有相对顺序。
	即："ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
	动规五步曲：
	1、dp数组和下标含义
		dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列为dp[i][j]
	2、递推公式
		如果text1[i - 1] 与 text2[j - 1]相同，
			那么找到了一个公共元素，所以dp[i][j] = dp[i - 1][j - 1] + 1;
		如果text1[i - 1] 与 text2[j - 1]不相同，
			那就看看text1[0, i - 2]与text2[0, j - 1]的最长公共子序列 和 text1[0, i - 1]与text2[0, j - 2]的最长公共子序列，取最大的。
		即：dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]);
	3、初始化：
		text1[0, i-1]和空串的最长公共子序列自然是0，所以dp[i][0] = 0;
		同理dp[0][j]也是0。
	4、遍历顺序
		要从前向后，从上到下来遍历这个矩阵。
	5、举例推导
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp数组和初始化
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	// 遍历
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}