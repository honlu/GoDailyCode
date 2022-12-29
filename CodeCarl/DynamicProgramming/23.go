/*
23. 打家劫舍2
2022-10-30
link: https://leetcode.cn/problems/house-robber-ii/
question:
	你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
	这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
	同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
	给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
answer:
	注意：多了一个条件，所有房屋围城一圈。即成环了。
	对于一个数组，成环的话，主要考虑如下三种情况：
	1、考虑不包含首尾元素。
	2、考虑包含首元素，不包含尾元素。
	3、考虑包含尾元素，不包含首元素。
	动态规划五步曲
*/

// 动态规划，时间和空间复杂度都是O(n)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 情况2下获得最大值
	result1 := robRange(nums, 0, len(nums)-2)
	// 情况3下获取最大值
	result2 := robRange(nums, 1, len(nums)-1)
	return max(result1, result2)
}

func robRange(nums []int, start, end int) int {
	if end == start { // 长度为2的数组
		return nums[start]
	}
	// 创建dp数组和初始化
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])
	// 遍历顺序
	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[end]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}