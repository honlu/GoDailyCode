package main

import "fmt"

func main() {
	t := 0
	fmt.Scanln(&t) // 换行时
	for k := 0; k < t; k++ {
		n, m := 0, 0
		fmt.Scan(&n, &m) // scanf用空格分隔，千万不能使用换行符，否则会都进去！因为使用缓存然后再读入。用Scan就可以绑定对应参数了！
		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			matrix[i] = make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Scan(&matrix[i][j])
			}
		}
		fmt.Println(matrix)
		res := 0
		fmt.Println(res)
	}
}
