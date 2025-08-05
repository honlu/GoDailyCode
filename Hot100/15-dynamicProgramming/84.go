package dynamicprogramming

import (
	"math"
)

/*
84-279.完全平方数
https://leetcode.cn/problems/perfect-squares/?envType=study-plan-v2&envId=top-100-liked
*/

/*
思路：个人想法
从最大平方数开始逆序计算。
 贪心剪枝 + DFS 回溯。
但要清楚：它并不能单独解决问题，因为：
贪心策略本身不足以保证最优解（类似「硬币找零」问题）
所以你需要结合 DFS 回溯 + 剪枝 + 记忆化等技术

非最优解
*/
// func numSquares(n int) int {
// 	// 预处理所有 <= n 的完全平方数
// 	squares := []int{}
// 	for i := 1; i*i <= n; i++ {
// 		squares = append(squares, i*i)
// 	}

// 	// 倒序，从大的平方数开始尝试
// 	for i, j := 0, len(squares)-1; i < j; i, j = i+1, j-1 {
// 		squares[i], squares[j] = squares[j], squares[i]
// 	}

// 	minCount := math.MaxInt32

// 	var dfs func(remain int, count int)
// 	dfs = func(remain int, count int) {
// 		if remain == 0 {
// 			if count < minCount {
// 				minCount = count
// 			}
// 			return
// 		}
// 		if count >= minCount {
// 			return // 剪枝
// 		}
// 		for _, square := range squares {
// 			if square > remain {
// 				continue
// 			}
// 			dfs(remain-square, count+1)
// 		}
// 	}

// 	dfs(n, 0)
// 	return minCount
// }

/*
思路：
参考答案的最优化动态规划解：外循环必须从小到大
*/
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ { // 外循环一定从小到大
		dp[i] = i                                          // 最坏情况：1+1+..+1，共i个
		for j := int(math.Sqrt(float64(i))); j >= 1; j-- { // 内循环，可以从大到小枚举平方数，也可从小到大
			dp[i] = min(dp[i], dp[i-j*j]+1)
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
