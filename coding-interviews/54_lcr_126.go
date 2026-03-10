package main

// 斐波那契数
func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i < n; i++ { // 动态规划
		a, b = b, a+b
	}
	return b
}
