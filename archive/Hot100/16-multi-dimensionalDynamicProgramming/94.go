package multidimensionaldynamicprogramming

/*
94-1143. 最长公共子序列
https://leetcode.cn/problems/longest-common-subsequence/?envType=study-plan-v2&envId=top-100-liked
*/
/*
动态规划：
dp[i][j]=在 text1 前 i 个字符-text1[0:i]和 text2 前 j 个字符-text2[0:j]中LCS 的长度，注意切片右边界为开区间

*/
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
