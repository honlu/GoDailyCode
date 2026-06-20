package lcr180

import "math"

// LCR 180 : 文件组合
// 注意：需要连续正整数
func fileCombination(target int) [][]int {
	var res [][]int
	i, j := 1.0, 2.0
	for i < j {
		// 数学公式，一元二次方式的解
		j = (-1 + (math.Sqrt(1 + 4*(2*float64(target)+i*i-i)))) / 2
		if i < j && j == float64(int(j)) { // j要为正整数的float64
			var temp []int
			for k := int(i); k <= int(j); k++ {
				temp = append(temp, k)
			}
			res = append(res, temp)
		}
		i++
	}
	return res
}
