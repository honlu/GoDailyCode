/*
8. 最长重复子数组
2022-11-2
link:https://leetcode.cn/problems/maximum-length-of-repeated-subarray/
question:
	给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
answer:
	注意题目中说的子数组，其实就是连续子序列。这种问题动规最拿手.
	动规五步曲：
	1、dp数组和下标含义[关键]：
		dp[i][j] ：以下标i - 1为结尾的A，和以下标j - 1为结尾的B，最长重复子数组长度为dp[i][j]。
		（特别注意： “以下标i - 1为结尾的A” 标明一定是 以A[i-1]为结尾的字符串 ）
		在遍历dp[i][j]的时候i 和 j都要从1开始。
	2、递推公式：
		当A[i - 1] 和B[j - 1]相等的时候，dp[i][j] = dp[i - 1][j - 1] + 1;
	3、初始化：
		根据dp[i][j]的定义，dp[i][0] 和dp[0][j]其实都是没有意义的！
		但dp[i][0] 和dp[0][j]要初始值，因为 为了方便递归公式dp[i][j] = dp[i - 1][j - 1] + 1;
		所以dp[i][0] 和dp[0][j]初始化为0。
	4、遍历顺序：
		外层for循环遍历A，内层for循环遍历B。
	5、举例推导
*/
func findLength(nums1 []int, nums2 []int) int {
	// dp数组和初始化
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(nums1)+1; i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	res := 0
	// 遍历方向
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}