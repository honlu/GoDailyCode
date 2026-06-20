package dynamicprogramming

/*
89-416.分割等和子集
https://leetcode.cn/problems/partition-equal-subset-sum/description/?envType=study-plan-v2&envId=top-100-liked
*/

/*
初步想法：总和要为2的倍数，才能分割。
然后在寻找分割对象，如果枚举出给一串数字，他们的合集是多少呢？
典型的「子集和=目标值」的问题
*/
func canPartition(nums []int) bool {
	var sum int
	for _, item := range nums {
		sum += item
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	// 0-1背包，用一维 DP 表示「是否能拼出 j」
	// dp[j] 表示是否可以用若干个数凑出和 j
	dp := make([]bool, target+1)
	dp[0] = true // base case: 和为 0 一定可以凑出（啥都不选）
	// 遍历每个元素（相当于物品）
	for _, item := range nums {
		// 倒序遍历容量，确保每个 item 只用一次
		// 如果正序，会导致同一元素被重复使用，变成完全背包
		for j := target; j >= item; j-- {
			// 转移方程：
			// 如果之前能凑出 j-item，那么加上当前 item 就能凑出 j
			dp[j] = dp[j] || dp[j-item]
		}
	}
	return dp[target]
}
