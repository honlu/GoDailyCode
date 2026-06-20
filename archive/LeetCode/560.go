/*
560. 和为k的子数组
2022-12-18
link: https://leetcode.cn/problems/subarray-sum-equals-k/
question:
	给你一个整数数组 nums 和一个整数 k ，
	请你统计并返回 该数组中和为 k 的连续子数组的个数 。
answer:
	注意：这里指的连续子数组，没有说离散情况。
	首先试试暴力解法，两层for循环.
	还有优化方法，待改进。
*/
// 暴力解法
func subarraySum(nums []int, k int) int {
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
		sum = 0
	}
	return res
}
