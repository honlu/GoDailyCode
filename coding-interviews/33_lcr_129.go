package main

// 字母迷宫 （单词搜索）
func wordPuzzle(grid [][]byte, target string) bool {
	rows := len(grid)
	cols := len(grid[0])
	// 定义DFS函数，参数为当前格子行列和目标单词的下标
	var dfs func(row, col, index int) bool
	dfs = func(row, col, index int) bool {
		// 如果已匹配到目标字符串末尾，返回true
		if index == len(target) {
			return true
		}
		// 越界检测
		if row >= rows || row < 0 || col >= cols || col < 0 {
			return false
		}
		// 已经访问或与目标字符不符，返回false
		if grid[row][col] == '0' || grid[row][col] != target[index] {
			return false
		}
		item := grid[row][col] // 记录当前格子的值，方便回溯
		grid[row][col] = '0'   // 标记当前格子已访问，进行剪枝
		// 尝试四个方向的搜索
		if dfs(row-1, col, index+1) || dfs(row+1, col, index+1) || dfs(row, col+1, index+1) || dfs(row, col-1, index+1) {
			return true
		} else {
			grid[row][col] = item // 回溯还原现场
		}
		return false
	}

	// 尝试每个起点
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 只从首字符相同的位置开始搜索
			if grid[i][j] == target[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
