/*
18、零钱兑换
2022-10-27
link: https://leetcode.cn/problems/coin-change/
question:
	给定不同面额的硬币 coins 和一个总金额 amount。
	编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
	如果没有任何一种硬币组合能组成总金额，返回 -1。
	你可以认为每种硬币的数量是无限的。
answer:
	动态规划五步曲：
	1、确定dp数组和下标含义：dp[i],凑够总额为i所需钱币的最少个数为dp[i]
	2、递推公式：dp[i] =min(dp[i-coins[j]+1, dp[i]]
	3、dp数组初始化：dp[0]=0,其他下标为最大值
	4、遍历顺序：两种循环都可以。建议target放在内循环，coins放在外循环
	5、举例推导验证
*/
func coinChange(coins []int, amount int) int {
	// 创建dp数组和初始化
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	// 遍历顺序
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}