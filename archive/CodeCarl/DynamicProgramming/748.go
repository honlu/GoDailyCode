/*
4
748. 使用最小花费爬楼梯
2022-10-14
update: 2023-2-23 by lu
(需要思考)
link: 748-https://leetcode.cn/problems/min-cost-climbing-stairs/
question:
	数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
	每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。
	请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。
answer:
	1、dp[i]的定义：到达第i个台阶所花费的最少体力为dp[i]。（注意这里认为是第一步一定是要花费）
	2、递推公式：dp[i] = min(dp[i - 1], dp[i - 2]) + cost[i];
	3、dp数组的初始化：dp[i]由dp[i-1]，dp[i-2]推出，既然初始化所有的dp[i]是不可能的，那么只初始化dp[0]和dp[1]就够了，其他的最终都是dp[0]dp[1]推出
	4、确定遍历顺序：因为是模拟台阶，而且dp[i]又dp[i-1]dp[i-2]推出，所以是从前到后遍历cost数组就可以
	5、草稿纸：举例推到dp数组是否正确
*/
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = dp[i-2]
		if dp[i] > dp[i-1] { // 注意这里符号方向
			dp[i] = dp[i-1]
		}
		dp[i] += cost[i]
	}
	// 注意最后一步不用消费，所以取倒数第一步或第二步的最少值
	if dp[len(dp)-1] < dp[len(dp)-2] {
		return dp[len(dp)-1]
	}
	return dp[len(dp)-2]
}

// 不一样的版本
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}