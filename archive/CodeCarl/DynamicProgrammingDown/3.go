/*
3.买卖股票的最佳时机4
2022-11-2
link:https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
question:
	给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
	设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
answer:
	是动态规划：123.买卖股票的最佳时机III (opens new window)的进阶版，这里要求至多有k次交易。
	最大的区别就是这里要类比j为奇数是买，偶数是卖的状态。
	动规五步曲：
	1、dp数组和下标含义：
	使用二维数组 dp[i][j] ：第i天的状态为j，所剩下的最大现金是dp[i][j]
	j的状态表示为：
		0 表示不操作
		1 第一次买入
		2 第一次卖出
		3 第二次买入
		4 第二次卖出
		.....
	除了0以外，偶数就是卖出，奇数就是买入。
	题目要求是至多有K笔交易，那么j的范围就定义为 2 * k + 1 就可以了。
	2、递推公式
	dp[i][1] = max(dp[i - 1][0] - prices[i], dp[i - 1][1])
	dp[i][2] = max(dp[i - 1][1] + prices[i], dp[i - 1][2])
	3、初始化
	dp[0][0] = 0;dp[0][1] = -prices[0];dp[0][2] = 0;dp[0][3] = -prices[0];..
	同理可以推出dp[0][j]当j为奇数的时候都初始化为 -prices[0]
	4、遍历顺序: 从前往后
	5、举例推导
*/
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp数组和初始化
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}
	for i := 1; i < 2*k+1; i += 2 {
		dp[0][i] = -prices[0]
	}
	// 递推公式和遍历
	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k-1; j += 2 { // 注意这里！
			dp[i][j+1] = max(dp[i-1][j]-prices[i], dp[i-1][j+1])
			dp[i][j+2] = max(dp[i-1][j+1]+prices[i], dp[i-1][j+2])
		}
	}
	return dp[len(prices)-1][2*k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}