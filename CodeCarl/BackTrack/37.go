/*
16
37.解数独
day: 2022-10-9
update: 2023-2-6 by lu
link: https://leetcode.cn/problems/sudoku-solver/
question:
	编写一个程序，通过填充空格来解决数独问题
answer:
	棋盘搜索问题可以使用回溯法暴力搜索。
	注意：以下代码不是按照代码随想录
*/

func solveSudoku(board [][]byte) {
	backTrack(0, 0, board)
}

// 参数：行列，棋盘
func backTrack(i, j int, board [][]byte) bool {
	// 若有解，返回true; 无解，返回fasle
	// 终止条件
	if j == 9 {
		// 换到下一行
		i++
		j = 0
		if i == 9 { // 棋盘已经填满，退出
			return true
		}
	}
	if board[i][j] != '.' { // 这里重要哦！
		return backTrack(i, j+1, board)
	}
	// (i,j) 这个位置放k是否合适
	for k := '1'; k <= '9'; k++ { // 注意这里小于等于号
		if isValid(i, j, byte(k), board) == false {
			continue
		}
		board[i][j] = byte(k) // 满足就填入
		// 递归
		if backTrack(i, j+1, board) {
			return true
		}
		// 回溯
		board[i][j] = '.'
	}
	// 若数字1-9都不能成功填入空格，返回false无解
	return false
}

/**
 * 判断在（row,col）插入k后，棋盘是否合法有如下三个维度:
 *     同行是否重复
 *     同列是否重复
 *     9宫格里是否重复
 */
func isValid(row, col int, k byte, board [][]byte) bool {
	// 同行是否重复 同列是否重复
	for i := 0; i < 9; i++ {
		if board[row][i] == k || board[i][col] == k {
			return false
		}
	}
	// 9宫格是否重复
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}