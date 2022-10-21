/*
动态规划：关于01背包问题！
对于面试，掌握01背包和完全背包，就够用了。最多再来一个多重背包。！
【重点：01背包，完全背包】
01背包理论基础1[二维dp数组]
2022-10-21

01背包原理：
问题：有n件物品和一个最多能背重量为w 的背包。第i件物品的重量是weight[i]，
	得到的价值是value[i] 。每件物品只能用一次，求解将哪些物品装入背包里物品价值总和最大。
解法：
	暴力解法：指数级的时间复杂度。
	每一件物品其实只有两个状态，取或者不取，所以可以使用回溯法搜索出所有的情况，那么时间复杂度就是o(2^n)，这里的n表示物品数量。
	然后动态规划进行优化。

	二位DP数组01背包：五部曲
	1、确定dp数组以及下标含义：dp[i][j] 表示从下标为[0-i]的物品里任意取，放进容量为j的背包，价值总和最大是多少。
	2、递推公式：dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i]);
	3、dp数组初始化[重难点]：
		首先从dp[i][j]的定义出发，如果背包容量j为0的话，即dp[i][0]，无论是选取哪些物品，背包价值总和一定为0。
		由状态转移方程可以看出i 是由 i-1 推导出来，那么i为0的时候就一定要初始化。
		其他下标初始为什么数值都可以，因为都会被覆盖。初始-1，初始-2，初始100，都可以！但只不过一开始就统一把dp数组统一初始为0，更方便一些。
	4、确定遍历顺序[重难点]：有两个遍历的维度：物品与背包重量。
		先遍历 物品还是先遍历背包重量呢？其实都可以！！ 但是先遍历物品更好理解。
	5、举例推到dp数组。
*/
package main

import "fmt"

func main() {
	weight := []int{1, 3, 4}   // 背包重量
	value := []int{15, 20, 30} // 背包value价值
	bagWeight := 4
	res := testWeightBagProblem(weight, value, bagWeight)
	fmt.Println(res)
}

func testWeightBagProblem(weight, value []int, bagWeight int) int {
	// 1.定义二维dp数组
	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, bagWeight+1) // 为什么初始化这个大小呢？
	}
	// 2.初始化
	for j := bagWeight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}
	// 3.递推公式
	for i := 1; i < len(weight); i++ {
		// 正序，也可以倒序
		for j := 0; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j] // 注意这里
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i]) // 注意这里！
			}
		}
	}
	return dp[len(weight)-1][bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
