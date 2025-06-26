package binarysearch

/*
64-搜索二维矩阵
https://leetcode.cn/problems/search-a-2d-matrix/description/?envType=study-plan-v2&envId=top-100-liked
*/
func searchMatrix(matrix [][]int, target int) bool {
	// 从右上角看是一个二叉搜索树
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target < matrix[i][j] {
			j--
		} else {
			i++
		}
	}
	return false
}
