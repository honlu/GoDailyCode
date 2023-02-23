/*
4
53. 最大子序和
2022-10-12
update: 2023-2-9 by lu
link：https://leetcode.cn/problems/maximum-subarray/
question：
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组
（子数组最少包含一个元素），返回其最大和
answer:
	暴力解法：两层for循环，第一层设置起始位置，第二层遍历数组寻找最大值
	注意：Go的暴力会超时，C++可以勉强通过. 为什么呢？

	贪心解法
*/
package main

import "math"

// 贪心解法—— 选择局部最优解
func maxSubArray(nums []int) int {
	res := math.MinInt
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > res { // 区间终止：其实就是如果count取到最大值了，及时记录
			res = count // 取区间累计的最大值（即不断确定最大子序终止位置）
		}
		if count <= 0 { // 当前“连续和”为负数的时候立刻放弃，从下一个元素重新计算“连续和”
			count = 0 // 重置最大子序起始位置，因为遇到附属一定时拉低总和
		}
	}
	return res
}

// 动态规划- 在原函数上进行动态规划的求解方法
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

// 暴力解法
// func maxSubArray(nums []int) int {
// 	result := math.MinInt
// 	count := 0
// 	for i := 0; i < len(nums); i++ {
// 		count = 0
// 		for j := i; j < len(nums); j++ { // 注意j要从i开始
// 			count += nums[j]
// 			if count > result {
// 				result = count
// 			}
// 		}
// 	}
// 	return result
// }
