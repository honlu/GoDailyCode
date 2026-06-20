/*
12、目标和
2022-10-23
link: https://leetcode.cn/problems/target-sum/
question:
	给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
	返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
answer:
	解法1：回溯法，暴力搜索出来，但会超时。
	解法2：转化为动态规划背包问题。
	关键点：如何使表达式结果为target，
	分析：
	既然为target，那么就一定有 left组合 - right组合 = target。
	left + right等于sum，而sum是固定的。
	公式来了， left - (sum - left) = target -> left = (target + sum)/2 。
	target是固定的，sum是固定的，left就可以求出来。
	此时问题就是在集合nums中找出和为left的组合。

	1、确定dp数组和下标含义：
	dp[j] 表示：填满j（包括j）这么大容积的包，有dp[j]种方法.
	2、确定递推公式：dp[j] += dp[j - nums[i]]
	不考虑nums[i]的情况下，填满容量为j的背包，有dp[j]种方法。
	那么考虑nums[i]的话（只要搞到nums[i]），凑成dp[j]就有dp[j - nums[i]] 种方法。
	3、DP数组初始化：
	递归公式可以看出，在初始化的时候dp[0] 一定要初始化为1。其他下标对应的数值应该初始化为0。
	因为dp[0]是在公式中一切递推结果的起源，如果dp[0]是0的话，递归结果将都是0。
	dp[j]要保证是0的初始值，才能正确的由dp[j - nums[i]]推导出来。
	4、确定遍历顺序
	前面讲过对于01背包问题一维dp的遍历，nums放在外循环，target在内循环，且内循环倒序。
	5、推导DP数组。
*/
// 时间复杂度：O(n × m)，n为正数个数，m为背包容量
// 空间复杂度：O(m)，m为背包容量
package main

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if math.Abs(target).(int) > sum || (target+sum)%2 == 1 { // 此时没有方案
		return 0
	}
	bagSize := (target + sum) / 2 // 背包大小
	if bagSize < 0 {
		return 0
	}
	// dp数组创建和初始化
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}
