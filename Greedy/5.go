/*
5、买卖股票的最佳时机
2022-10-12
link:122-https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
question:
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易。
answer:
	注意：只有一只股票；当前只能买或卖操作，不能同时参与多笔交易。
	局部最优：收集每天的正利润，全局最优：求得最大利润。
	局部最优可以推出全局最优，找不出反例，试一试贪心！
*/

// 贪心算法
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 { // 只要i天比前一天高，就可以在i-1天买入，i天卖出赚取利润
			res += prices[i] - prices[i-1]
		}
	}
	return res
}