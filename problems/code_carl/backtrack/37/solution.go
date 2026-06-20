package backtrack37

func solveSudoku(board [][]byte) {
	var backtrack func() bool
	backtrack = func() bool {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] != '.' {
					continue
				}
				for val := byte('1'); val <= '9'; val++ {
					if !isValid(row, col, val, board) {
						continue
					}
					board[row][col] = val
					if backtrack() {
						return true
					}
					board[row][col] = '.'
				}
				return false
			}
		}
		return true
	}
	backtrack()
}

func isValid(row, col int, val byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == val || board[i][col] == val {
			return false
		}
	}
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}
