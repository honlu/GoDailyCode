package backtrack

import "strings"

/*
62-51.N皇后
*/
func solveNQueens(n int) [][]string {
	var res [][]string
	// 用一个二维矩阵保存棋盘，先初始化
	matrix := make([][]string, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]string, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = "." // 初始都是空位
		}
	}
	var dfs func(row int) // 参数为第多少行的n皇后
	dfs = func(row int) {
		if row == n {
			// 表示前n行皇后位置都确定了
			var temp []string
			for _, item := range matrix {
				itemStr := strings.Join(item, "")
				temp = append(temp, itemStr)
			}
			res = append(res, temp)

		}
		// 判断当前行所有位置，满足皇后的情况，一行只能有一个位置放皇后
		for i := 0; i < n; i++ {
			if isValid(row, i, matrix) {
				matrix[row][i] = "Q"
				dfs(row + 1)         // 递归，下一行
				matrix[row][i] = "." // 回溯
			}
		}
	}

	// 开始执行第一行的每个可能
	dfs(0)

	return res
}

// 判断（row,column）在棋盘中的位置是否适合放皇后
// 条件: 相同列，同一斜线上不能已有皇后
func isValid(row, column int, matrix [][]string) bool {
	for i := 0; i < row; i++ { // 相同列
		if matrix[i][column] == "Q" {
			return false
		}
	}
	for i, j := row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 { // 左上对角线
		if matrix[i][j] == "Q" {
			return false
		}
	}
	// 右上对角线
	for i, j := row-1, column+1; i >= 0 && j < len(matrix); i, j = i-1, j+1 {
		if matrix[i][j] == "Q" {
			return false
		}
	}
	return true
}
