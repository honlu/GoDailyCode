package hot100

/*
20-旋转图像
n*n的二维矩阵顺时针旋转90度，第一列变成第一行倒叙，以此类推
*/
func rotate(matrix [][]int) {
	n := len(matrix) - 1
	for i := 0; i <= n; i++ {
		for j := i; j <= n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n/2; j++ {
			matrix[i][j], matrix[i][n-j] = matrix[i][n-j], matrix[i][j]
		}
	}

}
