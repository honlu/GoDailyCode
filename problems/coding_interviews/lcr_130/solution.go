package lcr130

// WardrobeFinishing 返回可以整理的格子数。
func WardrobeFinishing(m int, n int, cnt int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || digit(i)+digit(j) > cnt {
			return 0
		}
		visited[i][j] = true
		return 1 + dfs(i+1, j) + dfs(i, j+1)
	}
	return dfs(0, 0)
}

func digit(x int) int {
	var res int
	for x > 0 {
		res += x % 10
		x /= 10
	}
	return res
}
