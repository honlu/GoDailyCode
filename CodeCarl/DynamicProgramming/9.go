/*
01背包理论基础-一维dp数组[滚动数组]
2022-10-21

将二维dp将为一维dp，即滚动数组。
注意：背包问题的状态是可以压缩的。所以可以将二维dp数组转换。
发现：通过递推公式：dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i]);
其实可以发现如果把dp[i-1]那一层拷贝到dp[i]上，表达式完全可以是：
	dp[i][j] = max(dp[i][j], dp[i][j - weight[i]] + value[i]);
与其把dp[i - 1]这一层拷贝到dp[i]上，不如只用一个一维数组了，只用dp[j]（一维数组，也可以理解是一个滚动数组）。
滚动数组：满足上一层可以重复利用的条件，直接拷贝到当前层。

动态规划五部曲：
	1、确定dp数组：一维dp数组中，dp[j]表示：容量为j的背包，所背的物品价值可以最大为dp[j]。
	2、递推公式：dp[j]可以通过dp[j-weight[i]]推导出来，dp[j-weight[i]]表示容量为j-weight[i]的背包所背的最大价值。
	dp[j-weight[i]]+value[i]表示容量为 j-物品i重量 的背包加上 物品i的价值。
	（也就是容量为j的背包，放入物品i了之后的价值即：dp[j]）
	注意此时：dp[j]有两个选择，一个是取自己dp[j] 相当于 二维dp数组中的dp[i-1][j]，即不放物品i，
		一个是取dp[j-weight[i]]+value[i]，即放物品i，指定是取最大的，毕竟是求最大价值
	3、一维dp数组初始化：
		dp[j]表示：容量为j的背包，所背的物品价值可以最大为dp[j]，
			那么dp[0]就应该是0，因为背包容量为0所背的物品的最大价值就是0。
			那么非0下标都初始化为0就可以了。这样才能让dp数组在递归公式的过程中取的最大的价值，而不是被初始值覆盖了。
			[因为假设物品价值都是大于0的]
	4、一维dp数组遍历顺序：
		一维dp遍历的时候，背包是从大到小。倒序遍历是为了保证物品i只被放入一次！但如果一旦正序遍历了，那么物品0就会被重复加入多次！
	5、举例推导dp数组。
*/
package main

import "fmt"

func main() {
	weight := []int{1, 3, 4}   // 背包重量
	value := []int{15, 20, 30} // 背包value价值
	bagWeight := 4
	res := testWeightBagProblem2(weight, value, bagWeight)
	fmt.Println(res)
}

// 一维dp数组解决
func testWeightBagProblem2(weight, value []int, bagWeight int) int {
	// 定义一维dp数组。默认初始化全为0了。
	dp := make([]int, bagWeight+1)
	// 递推公式
	for i := 0; i < len(weight); i++ {
		// 这里必须倒叙。区别二维，因为二维dp保存了i的状态。
		for j := bagWeight; j >= weight[i]; j-- { // 注意这里倒序
			// 递推公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
