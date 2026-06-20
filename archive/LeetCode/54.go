/*
54. 螺旋矩阵
2022-11-17
link: https://leetcode.cn/problems/spiral-matrix/
question:
	给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，
	返回矩阵中的所有元素。
answer:
	关于矩阵的打印方式！
	递归或迭代模拟！[经典模拟题！]
	设置四个边界上下左右，然后只有两个对比：上和下的对比，左和右的对比

*/
func spiralOrder(matrix [][]int) []int {
	var res []int
	m, n := len(matrix), len(matrix[0]) // m下边界，n右边界
	if m == 0 || n == 0 {
		return res
	}
	c, r := 0, 0 // c左边界，r上边界
	for {
		// 上边界：从左向右移动到最右
		for i := c; i < n; i++ {
			res = append(res, matrix[r][i])
		}
		r++         // 重新设定上边界，若上边界大于下边界
		if r >= m { // 注意等于
			break
		}
		// 右边界：从上到下移动
		for i := r; i < m; i++ {
			res = append(res, matrix[i][n-1]) // 注意n-1
		}
		n--         // 重新设定右边界
		if n <= c { // 若右边界小于左边界(注意：不仅仅要小于，还要等于)
			break
		}
		// 下边界：从右到左
		for i := n; i > c; i-- {
			res = append(res, matrix[m-1][i-1]) // 注意m-1
		}
		m--         // 重新设定下边界
		if m <= r { // 注意等于
			break
		}
		// 左边界：从下到上
		for i := m; i > r; i-- {
			res = append(res, matrix[i-1][c])
		}
		c++
		if c >= n { // 注意等于
			break
		}
	}
	return res
}