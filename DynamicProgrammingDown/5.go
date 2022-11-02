/*
5.最佳买卖股票时机含手续费
2022-11-2
link: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
question:
	给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
	你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
	返回获得利润的最大值。
	注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
answer:
	贪心方法写过。
	相对于动态规划：122.买卖股票的最佳时机II (opens new window)，
	本题只需要在计算卖出操作的时候减去手续费就可以了，代码几乎是一样的。
	唯一差别在于递推公式部分
	动态规划五步曲：
	1、dp数组和初始化
		dp[i][0] 表示第i天持有股票所省最多现金。
		dp[i][1] 表示第i天不持有股票所得最多现金
	2、递推公式
		如果第i天持有股票即dp[i][0]， 那么可以由两个状态推出来
		第i-1天就持有股票，那么就保持现状，所得现金就是昨天持有股票的所得现金 即：dp[i - 1][0]
		第i天买入股票，所得现金就是昨天不持有股票的所得现金减去 今天的股票价格 即：dp[i - 1][1] - prices[i]
		所以：dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] - prices[i]);

		如果第i天不持有股票即dp[i][1]的情况， 依然可以由两个状态推出来
		第i-1天就不持有股票，那么就保持现状，所得现金就是昨天不持有股票的所得现金 即：dp[i - 1][1]
		第i天卖出股票，所得现金就是按照今天股票价格卖出后所得现金，注意这里需要有手续费了即：dp[i - 1][0] + prices[i] - fee
		所以：dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] + prices[i] - fee);
	3、初始化
	4、遍历顺序
	5、举例推导
*/
func maxProfit(prices []int, fee int) int {
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
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0]-fee) // 注意这里区别，手续费
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}