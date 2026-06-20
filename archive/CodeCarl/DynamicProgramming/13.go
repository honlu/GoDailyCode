/*
13、一和零
2022-10-25
link: https://leetcode.cn/problems/ones-and-zeroes/
question:
	给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
	请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
	子集概念：如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
answer:
	1、dp数组和下标的含义：dp[i][j]最多有i个0和j个1的strs的最大子集的大小为dp[i][j]。
	2、递推公式
	dp[i][j] 可以由前一个strs里的字符串推导出来，strs里的字符串有zeroNum个0，oneNum个1。
	dp[i][j] 就可以是 dp[i - zeroNum][j - oneNum] + 1。
	然后我们在遍历的过程中，取dp[i][j]的最大值。
	所以递推公式：dp[i][j] = max(dp[i][j], dp[i - zeroNum][j - oneNum] + 1);
	3、dp数组如何初始化
	01背包的dp数组初始化为0就可以。因为物品价值不会是负数，初始为0，保证递推的时候dp[i][j]不会被初始值覆盖。
	4、确定遍历顺序
	01背包一定是外层for循环遍历物品，内层for循环遍历背包容量且从后向前遍历！
	5、举例推导DP数组
*/

func findMaxForm(strs []string, m int, n int) int {
	// dp数组创建和初始化为0
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 遍历
	for _, val := range strs {
		oneNum, zeroNum := 0, 0
		for _, v := range val {
			if v == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}