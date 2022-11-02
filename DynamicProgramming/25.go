/*
25. 买卖股票的最佳时机
2022-11-2
link: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
question:
	给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
	你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
	返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
answer:
	暴力方法：两次for循环，寻找最优间距。
	贪心方法：因为股票就买卖一次，所以一次for循环，记录左边最小值，以及当前值与最小值的差值，记录最后的最大差值。
	动态规划：
	1、dp数组和下标含义：
		dp[i][0]表示第i天持有股票所得最多现金；
		dp[i][1]表示第i天不持有股票所得最多现金。
		注意：“持有”不代表就是当天“买入”！可能买入，也有可能是昨天就买入了，今天保持持有的状态。
		如果第i天持有股票即dp[i][0]， 那么可以由两个状态推出来:
		第i-1天就持有股票，那么就保持现状，所得现金就是昨天持有股票的所得现金 即：dp[i - 1][0]
		第i天买入股票，所得现金就是昨天不持有股票的所得现金减去 今天的股票价格 即：dp[i - 1][1] - prices[i]
	2、递推公式：
		dp[i][0] = max(dp[i - 1][0], -prices[i]);
		dp[i][1] = max(dp[i - 1][1], prices[i] + dp[i - 1][0]);
	3、初始化：
	dp[0][0]表示第0天持有股票，此时的持有股票就一定是买入股票了，因为不可能有前一天推出来，所以dp[0][0] -= prices[0];
	dp[0][1]表示第0天不持有股票，不持有股票那么现金就是0，所以dp[0][1] = 0;
	4、遍历顺序：从前往后遍历
	5、推理推导dp数组
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
		dp[i][0] = max(dp[i-1][0], -prices[i])
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