package hot100

/*
21-搜索二维矩阵2
https://leetcode.cn/problems/search-a-2d-matrix-ii/?envType=study-plan-v2&envId=top-100-liked
*/
func searchMatrix(matrix [][]int, target int) bool {
	// 直接暴力吧
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
