package dynamicprogramming

import "math"

/*
83-198：打家劫舍
https://leetcode.cn/problems/house-robber/?envType=study-plan-v2&envId=top-100-liked
*/
/*
思路：仍然是找规律，找公式，找好初始条件
*/
func rob(nums []int) int {
	// 记录到每家时的最大值，以及所有家中最大值
	// dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	dp := make([]int, len(nums))
	copy(dp, nums) // 深拷贝
	var res int = math.MinInt
	for i := 0; i < len(nums); i++ {
		if i == 1 { // 特殊情况，「2,1,1,2」
			dp[i] = max(dp[i], dp[i-1])
		} else if i >= 2 {
			dp[i] = max(dp[i-2]+dp[i], dp[i-1]) // 转换公式
		}
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
