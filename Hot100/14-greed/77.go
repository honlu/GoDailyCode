package greed

/*
77-买卖股票的最佳时机

贪心、动态规划。
如果两层for循环会出现bad case超时
*/
func maxProfit(prices []int) int {
	var max int
	var start int = prices[0] // 精简暴力两层for循环，一次就可以。start作为之前记录的股票最小价格
	for i := 1; i < len(prices); i++ {
		if prices[i] < start {
			start = prices[i]
		} else {
			if prices[i]-start > max {
				max = prices[i] - start
			}
		}
	}
	return max
}
