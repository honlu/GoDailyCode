package hot100

/*
18-矩阵置零
https://leetcode.cn/problems/set-matrix-zeroes/?envType=study-plan-v2&envId=top-100-liked
*/
// 哈希
func setZeroes(matrix [][]int) {
	rowMap := make(map[int]bool)
	columnMap := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rowMap[i] = true
				columnMap[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rowMap[i] {
				matrix[i][j] = 0
			}
			if columnMap[j] {
				matrix[i][j] = 0
			}
		}
	}
}
