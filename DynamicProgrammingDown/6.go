/*
6. 最长递增子序列
2022-11-2
link:https://leetcode.cn/problems/longest-increasing-subsequence/
question:
	给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
	子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
answer:
	最长上升子序列是动规的经典题目。
	动态规划五部分：
	1、dp数组和下标含义：dp[i]表示i之前包括i的以nums[i]结尾最长上升子序列的长度
	2、递推公式
		if (nums[i] > nums[j]) dp[i] = max(dp[i], dp[j] + 1);
		注意这里不是要dp[i] 与 dp[j] + 1进行比较，而是我们要取dp[j] + 1的最大值。
	3、初始化：每一个i，对应的dp[i]（即最长上升子序列）起始大小至少都是1
	4、遍历顺序：
		dp[i] 是有0到i-1各个位置的最长升序子序列 推导而来，那么遍历i一定是从前向后遍历。
		j其实就是0到i-1，遍历i的循环在外层，遍历j则在内层
	5、举例推导
*/

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// dp数组和初始化
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	res := 0
	// 遍历
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > res {
			res = dp[i] // 取最长的子序列数
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}