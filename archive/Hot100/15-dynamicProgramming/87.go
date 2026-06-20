package dynamicprogramming

/*
87-300.最长递增子序列
https://leetcode.cn/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=top-100-liked
*/

/*
1. dp入门O(n^2) DP
思路：
dp[i] 表示以 nums[i] 结尾的最长递增子序列长度。
转移方程：
dp[i] = max(dp[j] + 1)，其中 j < i && nums[j] < nums[i]。
答案是 max(dp[i])。
*/
func lengthOfLIS(nums []int) int {
	m := len(nums)
	var ans int = 1
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
