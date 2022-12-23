/*
1. 买卖股票的最佳时机2
2022-11-2
link: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
question:
	给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
	在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
	返回 你能获得的 最大 利润 。
answer:
	贪心算法前面已经讲解。
	动态规划：注意本题可以买卖多次

*/
// 版本1：时间复杂度O(n),空间复杂度O(n)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp数组和初始化
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	// 遍历和递推公式
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]) // 注意：这里是唯一区别
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
	}
	return dp[len(prices)-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}