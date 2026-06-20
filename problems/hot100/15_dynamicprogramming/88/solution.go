package p15dynamicprogramming88

/*
88-152.乘积最大子数组
https://leetcode.cn/problems/maximum-product-subarray/?envType=study-plan-v2&envId=top-100-liked
*/
/*
🔹 常见误区
很多人一上来就想：
我维护一个 dp[i]，表示以 i 结尾的最大乘积。
然后想写：
dp[i] = max(nums[i], dp[i-1]*nums[i])
⚠️ 但问题是 有负数存在。
比如 [-2, 3, -4]：
如果你只维护「最大」，会丢掉信息。
因为 -2 * 3 = -6（最小），但再乘上 -4 → 24（最大）。
也就是说：最小值有可能变成最大值。
所以必须同时维护「最大」和「最小」。
🔹 正确的 DP 思路
定义两个数组（或变量即可）：
maxF[i] = 以 i 结尾的最大乘积
minF[i] = 以 i 结尾的最小乘积
转移方程：
maxF[i] = max(nums[i], nums[i]*maxF[i-1], nums[i]*minF[i-1])
minF[i] = min(nums[i], nums[i]*maxF[i-1], nums[i]*minF[i-1])
因为 nums[i] 可能是正数、负数、零：
乘以正数 → 最大继续最大
乘以负数 → 最大和最小交换角色
乘以 0 → 重置
*/
func maxProduct(nums []int) int {
	// 看似很明显的动态规划，但是却不知道如何设计dp，可能简单想到设计一个dp
	// 实际要涉及两个dp
	maxF, minF := nums[0], nums[0]
	var res int = nums[0]
	for i := 1; i < len(nums); i++ {
		x, y := maxF, minF // 小心覆盖
		maxF = max(nums[i], max(x*nums[i], y*nums[i]))
		minF = min(nums[i], min(x*nums[i], y*nums[i]))
		res = max(res, maxF)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
