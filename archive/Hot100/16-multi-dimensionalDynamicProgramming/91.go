package multidimensionaldynamicprogramming

/*
91-62. 不同路径
https://leetcode.cn/problems/unique-paths/?envType=study-plan-v2&envId=top-100-liked
*/

/*
比较合适的动态规划：
二维DP：
dp[i][j] 表示到达 (i,j) 的路径数。
dp[i][j] = dp[i][j-1] + dp[i-1][j]
*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i > 1 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 1 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
