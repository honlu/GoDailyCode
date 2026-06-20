/*
36. 有效的数独
2022-12-31
link: https://leetcode.cn/problems/valid-sudoku/
question:
	请你判断一个 9 x 9 的数独是否有效。
	只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
		数字 1-9 在每一行只能出现一次。
		数字 1-9 在每一列只能出现一次。
		数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

answer:
	暴力试试：设置三个哈希表数组，行，列，矩阵.

*/
type Hash struct {
	hash map[int]bool
}

func isValidSudoku(board [][]byte) bool {
	var row [9]Hash
	var column [9]Hash
	var matrix [3][3]Hash
	for i := 0; i < 9; i++ {
		row[i].hash = make(map[int]bool) // map要初始化才能使用
		for j := 0; j < 9; j++ {
			column[j].hash = make(map[int]bool)
			matrix[i/3][j/3].hash = make(map[int]bool)
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] >= '0' && board[i][j] <= '9' { // 注意这里细节，等于号不能少。bug发现好久！
				temp := int(board[i][j] - '0')
				if row[i].hash[temp] == true {
					return false
				} else {
					row[i].hash[temp] = true
				}

				if column[j].hash[temp] == true {
					return false
				} else {
					column[j].hash[temp] = true
				}

				if matrix[i/3][j/3].hash[temp] == true {
					return false
				} else {
					matrix[i/3][j/3].hash[temp] = true
				}
			}

		}
	}

	return true
}