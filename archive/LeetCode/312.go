/*
312.戳气球
2022-11-4
link: https://leetcode.cn/problems/burst-balloons/
question:
	有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
	现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
	求所能获得硬币的最大数量。
answer:
	动态规划五步曲：
	1、dp数组和下标含义
	2、递推公式
	3、初始化
	4、遍历顺序
	5、举例推导
*/
func maxCoins(nums []int) int {
	n := len(nums)
	// dp创建和初始化
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	// 数组首位添加两个1
	temp := make([]int, n+2)
	temp[0], temp[n+1] = 1, 1
	for i := 1; i < n+1; i++ {
		temp[i] = nums[i-1]
	}
	// 遍历更新
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+temp[i]*temp[k]*temp[j])
			}
		}
	}
	return dp[0][n+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}