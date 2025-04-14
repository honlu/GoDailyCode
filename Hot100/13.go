package hot100

/*
13-最大子数组和
动态规划
时间复杂度O(n)，空间复杂度O(1)
动态规划的核心是找到状态转移方程
状态转移方程：dp[i] = max(dp[i-1]+nums[i], nums[i])
dp[i]表示以nums[i]结尾的最大子数组和
dp[i-1]+nums[i]表示以nums[i]结尾的子数组和
nums[i]表示以nums[i]为起点的子数组和
dp[i] = max(dp[i-1]+nums[i], nums[i])表示以nums[i]结尾的最大子数组和
*/
func maxSubArray(nums []int) int {
	// 当前i的最大子数组和max(dp[i-1]+nums[i],nums[i])
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
