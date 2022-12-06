/*
240. 搜索二维矩阵2
2022-12-6
link: https://leetcode.cn/problems/search-a-2d-matrix-ii/
question:
	编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。
	该矩阵具有以下特性：
	每行的元素从左到右升序排列。
	每列的元素从上到下升序排列。

answer:
	二分查找
*/

// 暴力解法
// func searchMatrix(matrix [][]int, target int) bool {
// 	for _, val := range matrix {
// 		for _, temp := range val {
// 			if temp == target {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// 二分查找，对每一行二分查找
func searchMatrix(matrix [][]int, target int) bool {
	for _, val := range matrix {
		temp := sort.SearchInts(val, target) // 返回要插入的坐标
		if temp < len(val) && val[temp] == target {
			return true
		}
	}
	return false
}