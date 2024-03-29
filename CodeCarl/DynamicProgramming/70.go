package dynamicprogramming

/*
3
70. 爬楼梯
day:2022-6-19
update: 2023-2-23 by lu [easy]
link:https://leetcode.cn/problems/climbing-stairs/
idea:
本题大家如果没有接触过的话，会感觉比较难，多举几个例子，就可以发现其规律。
爬到第一层楼梯有一种方法，爬到二层楼梯有两种方法。
那么第一层楼梯再跨两步就到第三层 ，第二层楼梯再跨一步就到第三层。
所以到第三层楼梯的状态可以由第二层楼梯 和 到第一层楼梯状态推导出来，那么就可以想到动态规划了。

然后开始，动态规划五部曲：
1、确定dp数组以及下标的含义 dp[i] 爬到第i层楼梯，有dp[i]中方法
2、确定递推公式 dp[i] = dp[i-1] + dp[i-2]
3、dp数组如何初始化 ：本体不讨论dp[0]的初始化，直接使用dp[1]=1, dp[2]=2方便理解
4、确定遍历顺序：从前往后
5、距离推到dp数组或者打印出来看看和推到是否一样
*/

// 动态规划解决爬楼梯问题
func climbStairs(n int) int {
	if n < 2 { // 避免特殊情况，n=1时，下面dp[2]超出下标
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2       // 初始化
	for i := 3; i <= n; i++ { // 从前往后遍历
		dp[i] = dp[i-2] + dp[i-1] // 递推公式
	}
	return dp[n]
}

// 不同起始位置更新
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
