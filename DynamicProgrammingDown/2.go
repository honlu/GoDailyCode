/*
2. 买卖股票的最佳时机3
2022-11-2
link: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
question:
	给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
	设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
answer:
	添加了一个注意事项。
	关键在于至多买卖两次，这意味着可以买卖一次，可以买卖两次，也可以不买卖。
	动态规划：
	1、dp数组和下标含义：
	一天一共就有五个状态:没有操作,第一次买入,第一次卖出,第二次买入,第二次卖出
	dp[i][j]中 i表示第i天，j为 [0 - 4] 五个状态，dp[i][j]表示第i天状态j所剩最大现金。
	2、递推公式：
	dp[i][1] = max(dp[i-1][0] - prices[i], dp[i - 1][1])
	dp[i][2] = max(dp[i - 1][1] + prices[i], dp[i - 1][2])
	dp[i][3] = max(dp[i - 1][3], dp[i - 1][2] - prices[i])
	dp[i][4] = max(dp[i - 1][4], dp[i - 1][3] + prices[i])
	3、初始化
	dp[0][0] = 0;dp[0][1] = -prices[0];dp[0][2] = 0;dp[0][3] = -prices[0];dp[0][4] = 0;
	4、遍历顺序
	从前往后遍历
	5、举例推导
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp数组和初始化
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	// 遍历和递推公式[关键]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}