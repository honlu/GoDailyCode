/*
4、最佳买卖股票时机含冷冻期
2022-11-2
link: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
question:
	给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。
	设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
	你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
answer:
	相对于动态规划：122.买卖股票的最佳时机II (opens new window)，本题加上了一个冷冻期
	在动态规划：122.买卖股票的最佳时机II (opens new window)中有两个状态，持有股票后的最多现金，和不持有股票的最多现金。
	动态规划五步曲：
	1、dp数组和下标含义
	dp[i][j]，第i天状态为j，所剩的最多现金为dp[i][j]。
	因为出现冷冻期之后，状态其实是比较复杂度，例如今天买入股票、今天卖出股票、今天是冷冻期，都是不能操作股票的。
	具体可以区分出如下四个状态：
	状态一：买入股票状态（今天买入股票，或者是之前就买入了股票然后没有操作）
	卖出股票状态，这里就有两种卖出股票状态:
		状态二：两天前就卖出了股票，度过了冷冻期，一直没操作，今天保持卖出股票状态
		状态三：今天卖出了股票
	状态四：今天为冷冻期状态，但冷冻期状态不可持续，只有一天！
	2、递推公式
		dp[i][0] = max(dp[i - 1][0], max(dp[i - 1][3], dp[i - 1][1]) - prices[i]);
		dp[i][1] = max(dp[i - 1][1], dp[i - 1][3]);
		dp[i][2] = dp[i - 1][0] + prices[i];
		dp[i][3] = dp[i - 1][2];
	3、初始化
		dp[0][0] = -prices[0]，买入股票所剩现金为负数。
		第0天没有卖出dp[0][1]初始化为0就行
		同样dp[0][2]初始化为0，因为最少收益就是0，绝不会是负数。
		同理dp[0][3]也初始为0。
	4、遍历顺序：从前往后遍历
	5、举例推导
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp数组和初始化
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	// 递推公式和遍历[关键]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3], dp[i-1][1])-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[len(prices)-1][3], max(dp[len(prices)-1][1], dp[len(prices)-1][2]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}