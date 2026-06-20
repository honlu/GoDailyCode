package multidimensionaldynamicprogramming

/*
95-72. 编辑距离
https://leetcode.cn/problems/edit-distance/?envType=study-plan-v2&envId=top-100-liked
*/

/*
动态规划：
转移方程，定义，初始化
定义：dp[i][j]将 word1 前 i 个字符转换为 word2 前 j 个字符的最少操作数.
状态转移
考虑 `word1[i-1]` 和 `word2[j-1]` 的最后一个字符：
(1) 如果相等:dp[i][j] = dp[i-1][j-1]
因为最后一个字符一样，不需要额外操作。
(2) 如果不相等,我们有三种选择，对应三种操作：
 1. 插入:     在 `word1[:i]` 后插入 `word2[j-1]`，相当于把 `dp[i][j-1]` 转过来，然后 +1。
    dp[i][j] = dp[i][j-1] + 1
 2. 删除:     删除 `word1[i-1]`，相当于从 `dp[i-1][j]` 转过来，然后 +1。
    dp[i][j] = dp[i-1][j] + 1
 3. 替换     把 `word1[i-1]` 替换成 `word2[j-1]`，相当于从 `dp[i-1][j-1]` 转过来，然后 +1。
    dp[i][j] = dp[i-1][j-1] + 1

取最小值：dp[i][j] = \min(

	dp[i-1][j] + 1, \quad   // 删除
	dp[i][j-1] + 1, \quad   // 插入
	dp[i-1][j-1] + 1        // 替换

)
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
