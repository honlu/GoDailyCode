/*
120. 三角形最小路径和
2023-3-30 by lu
link: https://leetcode.cn/problems/triangle/
question:
	给定一个三角形 triangle ，找出自顶向下的最小路径和。
	每一步只能移动到下一行中相邻的结点上。
	相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
	也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
answer:
	暴力：两层for循环，在triangle数组上更改相加，最后遍历最后一层找最小的路径和
*/
// 暴力
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	for i := 0; i < m; i++ {
		n := len(triangle[i])
		for j := 0; i > 0 && j < n; j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == n-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}
	// fmt.Println(triangle)
	res := math.MaxInt
	for i := 0; i < len(triangle[m-1]); i++ {
		if res > triangle[m-1][i] {
			res = triangle[m-1][i]
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}