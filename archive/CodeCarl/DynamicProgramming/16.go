/*
16、组合总和5
2022-10-27
link: https://leetcode.cn/problems/combination-sum-iv/
question:
	给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。
	请你从 nums 中找出并返回总和为 target 的元素组合的个数。
answer:
	注意：不限制数组中元素使用次数，所以是完全背包问题。
	本题描述说是求组合，但又说元素相同顺序不同的组合算两个组合，其实就是求排列！
	得到的集合是排列，就要考虑元素之间的遍历顺序。
	动态规划5步曲：
	1、dp数组和下标含义：dp[i],凑成目标正整数为i的排列个数为dp[i]
	2、递推公式：dp[i] += dp[i-nums[j]]
		dp[i]（考虑nums[j]）可以由 dp[i-nums[j]]（不考虑nums[j]）推导出来。
		因为只要得到nums[j]，排列个数dp[i-nums[j]]，就是dp[i]的一部分。
	3、初始化：dp[0]初始为1，其他dp[i]递归时才有数值基础。非0下标初始为0.
	4、遍历顺序：求排列就是外部for遍历背包target，内层for循环遍历物品nums。
	5、举例推导验证dp数组
*/
func combinationSum4(nums []int, target int) int {
	// 创建dp数组和初始化
	dp := make([]int, target+1)
	dp[0] = 1
	// 遍历
	for i := 0; i <= target; i++ { // 遍历背包
		for j := 0; j < len(nums); j++ { // 遍历物品
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}