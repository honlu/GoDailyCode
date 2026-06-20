/*
59. 螺旋矩阵2
2023-1-4
link: https://leetcode.cn/problems/spiral-matrix-ii/
question:
	给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，
	且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
answer:
	重点：矩阵的螺旋遍历
*/
package main

import "fmt"

func generateMatrix(n int) [][]int {
	left, right := 0, n
	up, down := 0, n
	matrix := make([][]int, n)
	num := 1
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for num <= n*n { // 条件不能出错，等于号不能少。或者写成left < right && up < down
		for i := left; i < right; i++ { // 从左到右
			matrix[up][i] = num
			num++
		}
		up++
		for i := up; i < down; i++ { // 从上到下
			matrix[i][right-1] = num
			num++
		}
		right--
		for i := right - 1; i >= left; i-- { // 从右到左 注意right-1，否则会覆盖
			matrix[down-1][i] = num
			num++
		}
		down--
		for i := down - 1; i >= up; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
