package main

import (
	"strings"
)

func dismantlingAction(arr string) int {
	// dp[i]代表连续不重复的招式数
	// dp[i]代表连续不重复的招式数
	count := len(arr)
	if count == 0 {
		return 0
	}
	dp := make([]int, count)
	dp[0] = 1
	res := 1
	for i := 1; i < count; i++ {
		start := i - dp[i-1]
		temp := arr[start:i]
		a := strings.Index(temp, string(arr[i]))
		if a >= 0 { // 注意等于0
			dp[i] = i - (a + start) // 注意这里 start+a
		} else {
			dp[i] = dp[i-1] + 1
		}
		// fmt.Println(a, temp, dp[i]) // 测试排查
		res = max(res, dp[i])
	}
	return res
}
