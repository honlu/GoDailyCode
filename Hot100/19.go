package hot100

/*
19-螺旋矩阵
https://leetcode.cn/problems/spiral-matrix/description/?envType=study-plan-v2&envId=top-100-liked

顺时针螺旋顺序打印矩阵
*/

func spiralOrder(matrix [][]int) []int {
	// 模拟操作，四个边界问题
	row, column := len(matrix)-1, len(matrix[0])-1
	i, j := 0, 0 // 临时行，列值
	var res []int
	for i <= row && j <= column {
		for temp := j; temp <= column; temp++ {
			res = append(res, matrix[i][temp])
		}
		i++
		if i > row {
			break
		}
		for temp := i; temp <= row; temp++ {
			res = append(res, matrix[temp][column])
		}
		column--
		if j > column {
			break
		}
		for temp := column; temp >= j; temp-- {
			res = append(res, matrix[row][temp])
		}
		row--
		if i > row {
			break
		}
		for temp := row; temp >= i; temp-- {
			res = append(res, matrix[temp][j])
		}
		j++
		if j > column { // 四种边界提前退出
			break
		}

	}
	return res
}
