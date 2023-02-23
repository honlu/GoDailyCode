/*
16
37.解数独
day: 2022-10-9
update: 2023-2-6 by lu
link: https://leetcode.cn/problems/sudoku-solver/
question:
	编写一个程序，通过填充9*9矩阵空格来解决数独问题。
	数独的解法需 遵循如下规则：
		数字 1-9 在每一行只能出现一次。
		数字 1-9 在每一列只能出现一次。
		数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
		数独部分空格内已填入了数字，空白格用 '.' 表示。
answer:
	提示：
	给定的数独序列只包含数字 1-9 和字符 '.' 。
	你可以假设给定的数独只有唯一解。
	给定数独永远是 9x9 形式的。

	代码随想录：二维递归，自创词汇，较难理解。

*/
// 回溯-代码随想录
func solveSudoku(board [][]byte) {
	// 回溯函数定义
	// 参数：隐藏的board, 注意有返回类型
	// 因为解数独找到一个符合的条件（就在树的叶子节点上）立刻就返回，
	// 相当于找从根节点到叶子节点一条唯一路径，所以需要使用bool返回值。
	var backTrack func() bool
	backTrack = func() {
		for i := 0; i < 9; i++ { // 遍历行
			for j := 0; j < 9; j++ { // 遍历列
				if board[i][j] == '.' {
					for k := '1'; k < '9'; k++ { // (i,j)位置放k是否合适
						if isValid(i, j, k, board) {
							board[i][j] = k  // 放置k
							if backTrack() { // 递归，如果找到合适的一组就返回
								return true
							}
							board[i][j] = '.' // 回溯
						}
					}
					return false // 9个数都试玩了，都不行，那么就返回false
				}
			}
		}
		return true // 遍历完，没有返回false，说明棋盘已经填充完毕
	}
	backTrack() // 调用回溯函数
}

// 回溯，函数提出来写
// 棋盘搜索问题可以使用回溯法暴力搜索。
// 	注意：以下代码不是按照代码随想录
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
	var k byte                   // 要先定义k为byte类型，否则go里面，'1'默认是rune类型保存
	for k = '1'; k <= '9'; k++ { // 注意这里小于等于号
		if isValid(i, j, k, board) == false {
			continue
		}
		board[i][j] = k // 满足就填入
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
		if board[row][i] == k || board[i][col] == k { // 同行或同列判断
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