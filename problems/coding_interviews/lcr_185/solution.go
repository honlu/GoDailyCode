package lcr185

// 感觉像是数学题
/*
概率是一样的，只是数字之和可能有重复的
只和数字范围[m~6*m]

从2为角度扩展到n维，解题思维牛：
- https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof/solutions/637778/jian-zhi-offer-60-n-ge-tou-zi-de-dian-sh-z36d/?envType=study-plan-v2&envId=coding-interviews

当一个随机过程增加一个独立状态时，枚举「旧状态 × 新状态」，把旧状态概率乘以新事件概率，累加到新状态。
「滚动数组 DP + 状态压缩」
*/
func statisticsProbability(num int) []float64 {
	// 1 个骰子有 6 种结果，每种结果概率都是 1/6
	// dp[j] 表示当前骰子数量为 i 时，
	// 点数和为「j + i」的概率（通过偏移让数组下标从 0 开始）
	dp := make([]float64, 6)
	for i := 0; i < 6; i++ {
		dp[i] = 1.0 / 6
	}
	// 从第 2 个骰子开始，逐个加入骰子进行状态转移
	for i := 2; i <= num; i++ {
		// i 个骰子的点数范围：[i, 6*i]
		// 一共有 6*i-i+1 = 5*i+1 种不同的点数和
		temp := make([]float64, 5*i+1)
		// 枚举前 i-1 个骰子的所有点数和
		for j := range len(dp) {
			// 新加入的骰子可能掷出 1~6
			// k=0~5 分别对应骰子点数 1~6
			for k := range 6 {
				// dp[j] 对应的实际点数和为：j + (i-1)
				// 新骰子点数为：k + 1
				//
				// 新的实际点数和：
				// j + (i-1) + (k+1) = j+k+i
				//
				// temp 下标需要减去最小点数和 i，
				// 所以新下标为：j+k
				//
				// 新骰子出现每个点数的概率都是 1/6，
				// 因此把 dp[j]/6 累加到对应的新状态
				temp[j+k] += dp[j] / 6
			}
		}
		// 滚动数组：temp 成为 i 个骰子的概率分布
		dp = temp
	}

	return dp
}
