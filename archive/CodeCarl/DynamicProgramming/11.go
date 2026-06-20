/*
11、最后一块石头的重量2
2022-10-22
link:https://leetcode.cn/problems/last-stone-weight-ii/
question:
	有一堆石头，每块石头的重量都是正整数。
	每一回合，从中选出任意两块石头，然后将它们一起粉碎。
	假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
	如果 x == y，那么两块石头都会被完全粉碎；
	如果 x != y，那么重量为x的石头将会完全粉碎，而重量为y的石头新重量为y-x。
	最后，最多只会剩下一块石头。
	返回此石头最小的可能重量。如果没有石头剩下，就返回 0。
answer:
	本题本质是：
		尽量让石头分成重量相同的两堆，相撞之后剩下的石头最小，这样就化解成01背包问题了。
	动态规划五步曲：
	1、确定dp数组和下标含义：
	dp[j]表示容量（这里说容量更形象，其实就是重量）为j的背包，最多可以背dp[j]这么重的石头。
	2、确定递推公式：dp[j] = max(dp[j], dp[j - stones[i]] + stones[i]);
	3、dp数组初始化：
		方法1：既然 dp[j]中的j表示容量，那么最大容量（重量）是多少呢，就是所有石头的重量和。
		因为提示中给出1 <= stones.length <= 30，1 <= stones[i] <= 1000，所以最大重量就是30 * 1000 。
		而我们要求的target其实只是最大重量的一半，所以dp数组开到15000大小就可以了。
		方法2：把石头遍历一遍，计算出石头总重量 然后除2，得到dp数组的大小。
	因为重量都不会是负数，所以dp[j]都初始化为0就可以了，这样在递归公式dp[j] = max(dp[j], dp[j - stones[i]] + stones[i]);
	中dp[j]才不会初始值所覆盖。
	4、确定遍历顺序：滚动数组，如果使用一维dp数组，物品遍历的for循环放在外层，遍历背包的for循环放在内层，且内层for循环倒序遍历！
	5、举例推导dp数组
*/
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, value := range stones {
		sum += value
	}
	target := sum / 2
	// dp数组确定和初始化
	dp := make([]int, target+1)
	for _, value := range stones { // 遍历物品即石头
		for j := target; j >= value; j-- { // 遍历背包
			dp[j] = max(dp[j], dp[j-value]+value)
		}
	}
	return sum - dp[target] - dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}