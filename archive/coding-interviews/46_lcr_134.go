package main

func myPow(x float64, n int) float64 {
	// 递归实现，主要区别n的正负数和零值
	if n == 0 {
		return 1
	} else if n < 0 {
		return myPow(1/x, -n)
	} else {
		// 如果直接：return x * myPow(x, n-1) 会出现超内存限制情况，所以为了减少内存使用，需要区分n的奇偶性
		if n%2 == 1 { // 奇数
			return x * myPow(x, n-1)
		} else { // 偶数
			sub := myPow(x, n/2) // 降低计算量
			return sub * sub
		}
	}
}
