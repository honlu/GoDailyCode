/*
7. 最长连续递增序列
2022-11-2
link: https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
question:
	给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
	连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，
	那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
answer:
	相对于昨天的最长递增子序列 (opens new window)最大的区别在于“连续”。
	动态规划五步曲：
	1、dp数组和下标含义[关键]
		dp[i]：以下标i为结尾的数组的连续递增的子序列长度为dp[i]。
		注意这里的定义，一定是以下标i为结尾，并不是说一定以下标0为起始位置。
	2、递推公式
		如果 nums[i + 1] > nums[i]那么dp[i + 1] = dp[i] + 1
	3、初始化：以下标i为结尾的数组的连续递增的子序列长度最少也应该是1，即就是nums[i]这一个元素。
		所以dp[i]应该初始1;
	4、遍历顺序
	 dp[i + 1]依赖dp[i]，所以一定是从前向后遍历。
	5、举例推导
*/
func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// dp数组和初始化
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	res := 1
	// 遍历
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}