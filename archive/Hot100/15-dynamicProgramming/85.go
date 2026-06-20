package dynamicprogramming

/*
85-322.零钱兑换
https://leetcode.cn/problems/coin-change/description/?envType=study-plan-v2&envId=top-100-liked
*/
/*
思路：
1. 暴力：自递归，即深度优先搜索，不是很好。
2. 广度有限搜索，动态规划，需要想好公式，初始值，每个值的含义，状态变化。
*/
func coinChange(coins []int, amount int) int {
	/*
		状态定义与转移公式
			状态定义：
			dp[i] 表示「凑出金额 i 所需的最少硬币数」
			初始化：
			dp[0] = 0（凑出 0 元，不需要任何硬币）
			其他 dp[i] = amount + 1（表示未被计算，或默认最大硬币数）
			状态转移公式：
			dp[i] = min(dp[i], dp[i - coin] + 1)
	*/
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // 最多
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = min(dp[i-coins[j]]+1, dp[i])
		}
	}
	if dp[amount] == amount+1 {
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
