package main

// 跳跃训练
func trainWays(num int) int {
	if num == 0 {
		return 1
	}
	if num <= 2 {
		return num
	}
	a, b := 1, 2
	for i := 3; i <= num; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}
