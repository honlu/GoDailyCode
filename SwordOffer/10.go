/*
剑指offer10-2. 青蛙跳台阶问题
2023-2-10  by lu
link: https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
question:
	一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。
	求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

	答案需要取模 1e9+7（1000000007），
	如计算初始结果为：1000000008，请返回 1
answer:
	爬到第 n 级台阶的方法个数等于爬到 n - 1 的方法个数和爬到 n - 2 的方法个数之和。
*/

// 动态规划
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	// 动态规划
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % (1e9 + 7)
	}
	return dp[n]
}

// 内存优化
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	// 优化内存使用
	a, b := 1, 1
	var res int
	for i := 2; i <= n; i++ {
		res = (a + b) % (1e9 + 7)
		a, b = b, res
	}
	return res
}