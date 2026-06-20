/*
5
122. 买卖股票的最佳时机
2022-10-12
update: 2023-2-10
link:122-https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
question:
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易。
answer:
	注意：
		只有一只股票；
		当前只能买或卖操作，不能同时参与多笔交易。
	贪心：
	局部最优：收集每天的正利润，全局最优：求得最大利润。
	局部最优可以推出全局最优，找不出反例，试一试贪心！
	这里有个特点：
		假如第0天买入，第3天卖出，那么利润为：prices[3] - prices[0]。
		相当于(prices[3] - prices[2]) + (prices[2] - prices[1]) + (prices[1] - prices[0])。
	此时就是把利润分解为每天为单位的维度，而不是从0天到第3天整体去考虑！


	动态规划更好理解：股票问题其实是一个系列的，属于动态规划的范畴
*/
/*
本题可以用贪心的关键：
1. 只有一只股票，只能买或卖操作
2. 类比这种正收入的操作。假如第0天买入，第3天卖出，那么利润为：prices[3] - prices[0]。
	等价于(prices[3] - prices[2]) + (prices[2] - prices[1]) + (prices[1] - prices[0])

当然本题用动态规划更好！
*/
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 { // 只要i天比前一天高，就可以在i-1天买入，i天卖出赚取利润
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

// 动态规划
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	// dp[i][0] 表示第i天持有股票所得最多现金。
	// dp[i][1] 表示第i天不持有股票所得最多现金
	dp[0][0] = 0 - prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		// 如果第i天还想持有股票即dp[i][0],可以由两种状态推出来
		// 第i-1天就持有，那么就保持现状，所得最多现金就是昨天持有股票的所得最多现金。即dp[i-1][0]
		// 第i-1天没有持有股票，第i天买入股票，所得现金就是昨天不持有股票的所得现金减去今天股票价格，即dp[i-1][1]-princes[i]
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-princes[i])

		// 如果第i天不持有股票即dp[i][1]的情况， 依然可以由两个状态推出来
		// 第i-1天就不持有股票，那么就保持现状，所得现金就是昨天不持有股票的所得现金 即：dp[i - 1][1]
		// 若第i-1天持有股票，第i天卖出股票，所得现金就是按照今天股票佳价格卖出后所得现金即：prices[i] + dp[i - 1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}