/*
18. 买卖股票地最佳时机含手续费
2022-10-13
link:714-https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
question:
	给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；
	非负整数 fee 代表了交易股票的手续费用。可以无限次地完成交易，但是你每笔交易都需要付手续费。
	如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
	求获得利润的最大值。
answer:
	和5.go 买卖股票地最佳时机类似，多添加了一个手续费。
	贪心策略，就是最低值买，最高值（如果算上手续费还盈利）就卖。

	做收获利润操作的时候其实有三种情况：
	情况一：收获利润的这一天并不是收获利润区间里的最后一天（不是真正的卖出，相当于持有股票），所以后面要继续收获利润。
	情况二：前一天是收获利润区间里的最后一天（相当于真正的卖出了），今天要重新记录最小价格了。
	情况三：不作操作，保持原有状态（买入，卖出，不买不卖）[代码实现的部分可以忽略]
*/

func maxProfit(prices []int, fee int) int {
	res := 0
	min := prices[0] // 记录最低价格
	for i := 1; i < len(prices); i++ {
		// 情况2，相当于买入
		if prices[i] < min {
			min = prices[i] // 此时买入
		}
		// 情况3，保持原有状态（因为此时买不便宜，卖则亏本）
		if prices[i] >= min && prices[i] <= min+fee {
			continue
		}
		// 计算利润，可能有多次计算利润，最后一次计算利润才是真正意义的卖出
		// 这里有些绕，很关键！
		if prices[i] > min+fee {
			res += prices[i] - min - fee
			min = prices[i] - fee // 情况1，这一步很关键！不好懂哦！要多思考和举例
		}
	}
	return res
}