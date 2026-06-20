package lcr163

import (
	"strconv"
)

// 数学规律
func findKthNumber(k int) int {
	digit, start, count := 1, 1, 9 // 1位数，从1开始，有9个
	for k > count {
		k = k - count
		digit += 1
		start *= 10
		count = 9 * digit * start
	}
	num := (k-1)/digit + start // 定位对应的数字
	numStr := strconv.Itoa(num)
	res, _ := strconv.Atoi(string(numStr[(k-1)%digit])) // 定位数字的哪一位
	return res
}
