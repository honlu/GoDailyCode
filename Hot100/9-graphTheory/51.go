package hot100

/*
51-岛屿数量
https://leetcode.cn/problems/number-of-islands/description/?envType=study-plan-v2&envId=top-100-liked

经典的flood fill 算法题。
*/
func numIslands(grid [][]byte) int {
	// 深度优先遍历或广度优先遍历
	// 深度优先遍历
	m, n := len(grid), len(grid[0])
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	var res [][]byte
	var cnt int
	res = grid
	var dfs func(x, y int)
	dfs = func(x, y int) {
		res[x][y] = '0'
		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a >= 0 && a < m && b >= 0 && b < n && res[a][b] == '1' {
				dfs(a, b)
			}
		}
	}
	// 套路模板
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if res[i][j] == '1' {
				dfs(i, j)
				cnt++
			}
		}
	}
	return cnt
}
