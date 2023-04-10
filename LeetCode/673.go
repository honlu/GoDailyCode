/*
673. 最长递增子序列的个数
2023-4-10 by lu
link: https://leetcode.cn/problems/number-of-longest-increasing-subsequence/
question:
	给定一个未排序的整数数组 nums，返回最长递增子序列的个数 。
	注意 这个数列必须是 严格 递增的。
answer:
	如何获得最长递增子序列呢？
	如何边获得便统计最长递增子序列的个数呢？
	有点难，参考题解。
	利用动态规划！关于动态规划还是理解不够透彻！
*/
func findNumberOfLIS(nums []int) int {
	n, maxLen, res := len(nums), 0, 0
	dp := make([]int, n)  // dp[i]:以nums[i]结尾的最长序列长度
	cnt := make([]int, n) // cnt[i]:以nums[i]结尾的最长序列长度的个数
	for i, x := range nums {
		dp[i], cnt[i] = 1, 1 // 初始化很重要，说明是否理解思路
		for j, y := range nums[:i] {
			if x > y {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[i] + 1
					cnt[i] = cnt[j] // 重置计数
				} else if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j] // ? cnt[i]等于所有满足dp[j]+1=dp[i]的cnt[j]之和
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			res = cnt[i] // 重置计数
		} else if dp[i] == maxLen {
			res += cnt[i]
		}
	}
	return res
}