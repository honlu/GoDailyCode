package lcr131

// 砍竹子1
/*
主要要写好动态公式：
当1≤j<i，dp[i]=max {max(j×(i−j),j×dp[i−j])}
*/
func cuttingBamboo(bamboo_len int) int {
	if bamboo_len < 2 {
		return 0
	}
	if bamboo_len == 2 {
		return 1
	}
	if bamboo_len == 3 {
		return 2
	}
	// 动态规划
	dp := make([]int, bamboo_len+1)
	dp[2], dp[3] = 1, 2
	for i := 4; i <= bamboo_len; i++ {
		for j := 1; j < i; j++ {
			temp := max(j*(i-j), j*dp[i-j])
			if dp[i] < temp {
				dp[i] = temp
			}
		}
	}
	return dp[bamboo_len]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
