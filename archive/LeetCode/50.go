/*
50. Pow(x, n)
2023-1-18
link: https://leetcode.cn/problems/powx-n/
question:
	实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
answer：
	递归实现，n在满足一定情况下依次递减
*/
// 递归思想解决
func myPow(x float64, n int) float64 {
	// 特殊情况处理
	if n == 0 {
		return 1
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	// 正常情况处理即n>0
	if n%2 == 1 { // 奇数
		return x * myPow(x, n-1)
	} else { // 偶数
		sub := myPow(x, n/2) // 降低计算量
		return su * sub
	}
}

// 迭代实现(快速幂法)
// 向下整除 n/2 等价于右移一位 n>>1
// 去余数n%2等价于判断二进制最右位n&1
func myPow(x float64, n int) float64 {
	// 特殊情况处理
	if n == 0 {
		return 1
	}
	if n < 0 {
		x, n = 1/x, -n
	}

	// 正常情况处理即n>0
	var res float64 = 1
	for n > 0 { // 等于0时跳出
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}