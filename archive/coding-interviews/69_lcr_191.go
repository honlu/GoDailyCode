package main

func statisticalResult(arrayA []int) []int {
	// 需要注意零值问题，边界情况
	// 所以需要左右边划分开
	count := len(arrayA)
	if count == 0 { // 注意这个边界问题
		return []int{}
	}
	res := make([]int, count)
	res[0] = 1
	for i := 1; i < count; i++ {
		res[i] = 1
		res[i] = res[i-1] * arrayA[i-1]
	}
	var temp int = 1
	for i := count - 2; i >= 0; i-- {
		temp *= arrayA[i+1]
		res[i] *= temp
	}
	return res
}
