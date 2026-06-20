package lcr129

// WordPuzzle 判断字母矩阵中是否存在目标路径。
func WordPuzzle(grid [][]byte, target string) bool {
	rows := len(grid)
	if rows == 0 || len(target) == 0 {
		return len(target) == 0
	}
	cols := len(grid[0])

	var dfs func(row, col, index int) bool
	dfs = func(row, col, index int) bool {
		if index == len(target) {
			return true
		}
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return false
		}
		if grid[row][col] == '0' || grid[row][col] != target[index] {
			return false
		}

		item := grid[row][col]
		grid[row][col] = '0'
		if dfs(row-1, col, index+1) || dfs(row+1, col, index+1) ||
			dfs(row, col+1, index+1) || dfs(row, col-1, index+1) {
			return true
		}
		grid[row][col] = item
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == target[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
