package main

import "strconv"

func crackNumber(ciphertext int) int {
	ciStr := strconv.Itoa(ciphertext)
	m := len(ciStr)
	if m < 2 {
		return 1
	}
	dp := make([]int, m)
	dp[0], dp[1] = 1, 1
	a := ciStr[:2]
	if a >= "10" && a <= "25" {
		dp[1] = 2
	}
	for i := 2; i < m; i++ {
		a := ciStr[i-1 : i+1]
		if a >= "10" && a <= "25" {
			dp[i] = dp[i-2] + dp[i-1] // 注意这里，模拟爬楼梯
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[m-1]
}
