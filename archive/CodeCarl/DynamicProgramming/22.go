/*
22. 打家劫舍
2022-10-29
link: https://leetcode.cn/problems/house-robber/
question:
	你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
	如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
answer:
	动态规划五步曲：
	1、确定dp数组和下标含义：dp[i],考虑下标i（包含i）以内的房屋，最多可以偷窃的金额为dp[i]
	2、确定递推公式：dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	3、dp数组初始化：dp[0] = nums[0], dp[1]= max(nums[0],nums[1]), 其他下标初始为0
	4、遍历顺序：从前往后遍历
	5、举例推导dp数组
*/
func rob(nums []int) int {
	n := len(nums)
	// 创建dp数组和初始化
	dp := make([]int, n) // 注意这里长度为n,不是n+1
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}