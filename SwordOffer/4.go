/*
剑指offer 4. 二维数组中的查找
2023-3-8
link: https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
question：
	在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，
	每一列都按照从上到下 非递减 的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

answer:
	观察特征，发现右上角可以看作二叉搜索的根节点！
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 { // 注意这个特殊情况！
		return false
	}
	// 利用二叉搜索树思想，从右上开始遍历
	n, m := 0, len(matrix[0])-1
	for n <= len(matrix)-1 && m >= 0 {
		if matrix[n][m] == target {
			return true
		} else if matrix[n][m] > target {
			m--
		} else {
			n++
		}
	}
	return false
}