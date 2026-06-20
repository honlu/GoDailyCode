//go:build undo
// +build undo

package main

// 以下：模拟，递归。结果：超时，需要优化。
func iceBreakingGame(num int, target int) int {
	temp := make([]int, num)
	for i := 0; i < num; i++ {
		temp[i] = i
	}
	var cur func([]int, int) int
	cur = func(arr []int, target int) int {
		count := len(arr)
		if count == 1 {
			return arr[0]
		}
		index := (target - 1) % count
		arr = append(arr[index+1:], arr[:index]...)
		return cur(arr, target)
	}
	return cur(temp, target)
}

// 约瑟夫环递推公式
