package main

import "fmt"

/*
2022-9-9 360笔试第2题 AC27% 因为go不能输入第二组数据
*/
func main() {
	t := 0
	fmt.Scanln(&t) // 换行时
	for k := 0; k < t; k++ {
		n, m := 0, 0
		fmt.Scanf("%d%d", &n, &m) // scan不能像Scanf指定类型
		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			matrix[i] = make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Scanf("%d", &matrix[i][j])
			}
		}
		// fmt.Println(matrix) // 测试
		// 寻找正方形区域
		maxSize := min(n, m)            // 正方形区域最大形状
		res := 0                        // 最终结果
		for i := 1; i <= maxSize; i++ { // 区域大小
			for x := 0; x <= n-i; x++ { // 注意区间
				for y := 0; y <= m-i; y++ { //(x,y)左上角起点
					sum := sumAri(matrix, x, y, i)
					res = max(res, sum)
				}
			}
		}
		fmt.Println(res)
	}
}

// 求当前区域的大小
func sumAri(matrix [][]int, x, y, size int) int {
	sum := 0
	// fmt.Println("--")               // 测试
	// fmt.Println(matrix, x, y, size) // 测试
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			if matrix[i][j] < 0 { // 如果当前区域有负数，直接返回0
				return 0
			}
			sum += matrix[i][j]
		}
	}
	// fmt.Println(sum)  // 测试
	// fmt.Println("**") // 测试
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
