/*
17、爬楼梯（进阶版）
2022-10-27
link:https://leetcode.cn/problems/climbing-stairs/
question:
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	[扩展：每次可以1，2,..直到m个台阶，问有多少种方法？上面的问题就是m=2]
answer:
	将题目改成完全背包的形式来解答。
	动态规划五步曲：
	1、确定dp数组和下标含义：dp[i],爬到有i个台阶的楼顶，有dp[i]种方法
	2、递推公式：dp[i] += dp[i-nums[j]]
	3、dp数组初始化：dp[0]=1,其他下标为0
	4、遍历顺序：求排列，则target放在外循环，nums放在内循环
	5、举例推导验证
*/
func climbStairs(n int) int {
	//定义dp数组和初始化
	dp := make([]int, n+1)
	dp[0] = 1
	// 本题物品只有两个1,2
	m := 2
	// 遍历顺序
	for j := 1; j <= n; j++ { //先遍历背包
		for i := 1; i <= m; i++ { //再遍历物品
			if j >= i {
				dp[j] += dp[j-i]
			}
			//fmt.Println(dp)
		}
	}
	return dp[n]
}