/*
11. 最大子序和
2022-11-2
link: https://leetcode.cn/problems/maximum-subarray/
question:
	给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
	子数组 是数组中的一个连续部分。
answer:
	贪心方法，前面讲过了。
	动态规划五步曲：
	1、dp数组和下标含义：
		·dp[i]：包括下标i之前的最大连续子序列和为dp[i]。
	2、递推公式
		dp[i] = max(dp[i - 1] + nums[i], nums[i]);
	3、初始化
		dp[0]=nums[0]
	4、遍历顺序:从前往后遍历
	5、举例推导
*/

// 动态规划
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// dp数组和初始化
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0] // 注意这个取值！是个技巧活
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