package dynamicprogramming21

func multiplePack(weight, value, nums []int, bagWeight int) int {
	expandedWeight := []int{}
	expandedValue := []int{}
	for i := range weight {
		for count := 0; count < nums[i]; count++ {
			expandedWeight = append(expandedWeight, weight[i])
			expandedValue = append(expandedValue, value[i])
		}
	}

	dp := make([]int, bagWeight+1)
	for i := range expandedWeight {
		for j := bagWeight; j >= expandedWeight[i]; j-- {
			dp[j] = max(dp[j], dp[j-expandedWeight[i]]+expandedValue[i])
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
