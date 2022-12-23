/*
19、完全平方数
2022-10-27
link: https://leetcode.cn/problems/perfect-squares/
question:
	给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
	完全平方数 是一个整数，其值等于另一个整数的平方；
	换句话说，其值等于一个整数自乘的积。
	例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
answer:
	转个问法：完全平方数就是物品（可以无限件使用），凑个正整数n就是背包，
	问凑满这个背包最少有多少物品？
	动态规划五步曲：
	1、确定dp数组和下标含义：dp[j],和为j的完全平方数的最少数量为dp[j]
	2、递推公式：dp[j] =min(dp[j-i*i]+1, dp[j]]
	3、dp数组初始化：dp[0]=0,其他下标为最大值
	4、遍历顺序：两种循环都可以。
		建议外循环遍历背包，内循环遍历物品
	5、举例推导验证
*/
func numSquares(n int) int {
	// 创建dp和初始化
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	// 遍历顺序
	for i := 0; i <= n; i++ { // 遍历背包
		for j := 1; j*j <= i; j++ { // 遍历物品
			dp[i] = min(dp[i-j*j]+1, dp[i]) // 关键
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}