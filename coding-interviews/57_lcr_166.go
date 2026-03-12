package main

func jewelleryValue(frame [][]int) int {
	// 找公式
	m, n := len(frame), len(frame[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var a, b int
			if i-1 >= 0 {
				a = frame[i-1][j]
			}
			if j-1 >= 0 {
				b = frame[i][j-1]
			}
			frame[i][j] += max(a, b)
		}
	}
	return frame[m-1][n-1]
}
