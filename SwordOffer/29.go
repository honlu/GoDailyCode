/*
剑指29. 顺时针打印矩阵
2023-3-22
link: https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
question:
	输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
answer:
	遍历模拟
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	up, down := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	var res []int
	for up <= down && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// 这里必须加
		if up > down || left > right {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}