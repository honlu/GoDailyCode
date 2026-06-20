package dynamicprogramming

/*
2
509. 斐波那契数
day:2022-6-17
update: 2023-2-23 by lu
link:https://leetcode.cn/problems/fibonacci-number/
idea:
「代码随想录」的风格是：简单题目是用来加深对解题方法论的理解的。
通过这道题目让大家可以初步认识到，按照动规五部曲是如何解题的。
对于动规，如果没有方法论的话，可能简单题目可以顺手一写就过，难一点就不知道如何下手了。
1-确定dp数组以及下标的含义
dp[i]的定义为：第i个数的斐波那契数值是dp[i]
2-确定递推公式
因为题目已经把递推公式直接给我们了：状态转移方程 dp[i] = dp[i - 1] + dp[i - 2];
3-dp数组如何初始化
题目中把如何初始化也直接给我们了，如下：
dp[0] = 0;
dp[1] = 1;
4-确定遍历顺序
从递归公式dp[i] = dp[i - 1] + dp[i - 2];中可以看出，dp[i]是依赖 dp[i - 1] 和 dp[i - 2]，那么遍历的顺序一定是从前到后遍历的
5-举例推导dp数组
按照这个递推公式dp[i] = dp[i - 1] + dp[i - 2]，我们来推导一下，当N为10的时候，dp数组应该是如下的数列：
0 1 1 2 3 5 8 13 21 34 55
如果代码写出来，发现结果不对，就把dp数组打印出来看看和我们推导的数列是不是一致的。
*/

// 动态规划实现（只用两个长度的数组+一个结果变量，不断互换更新）
func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := [2]int{0, 1}
	var sum int
	for i := 2; i <= n; i++ {
		sum = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return sum
}

// 数组完成版
func fib(n int) int {
	if n <= 1 {
		return n
	}
	// 创建dp数组，理解dp每个i的含义
	var dp []int
	dp = make([]int, n+1)
	// 确定dp递推公式；初始化；遍历更新
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
