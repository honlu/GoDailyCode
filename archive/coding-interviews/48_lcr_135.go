package main

import "math"

func countNumbers(cnt int) []int {
	// 最大正整数10的cnt次方-1
	end := int(math.Pow(10.0, float64(cnt)*1.0))
	res := make([]int, 0, end+1) // 注意这里要预分配空间
	for i := 1; i < end; i++ {
		res = append(res, i)
	}
	return res
}
