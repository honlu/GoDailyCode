package dynamicprogramming14

func completePack(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= bagWeight; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
