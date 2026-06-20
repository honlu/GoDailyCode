package dynamicprogramming12

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if target < -sum || target > sum || (target+sum)%2 != 0 {
		return 0
	}

	bagSize := (target + sum) / 2
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for _, num := range nums {
		for j := bagSize; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[bagSize]
}
