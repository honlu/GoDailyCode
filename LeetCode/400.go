/*
400. 第N位数字
2023-3-27 by lu
link: https://leetcode.cn/problems/nth-digit/
question:
	给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...] 中找出并返回第 n 位上的数字。
answer:
	注意模拟的计算公式。
	n位长度对应范围数：9*1e(n-1)*n
*/
func findNthDigit(n int) int {
	a, b := 9, 1 // b长度的范围，b数字的长度
	for n > a*b {
		n -= a * b
		a *= 10
		b++
	}
	res := strconv.Itoa((n-1)/b + a/9)
	// fmt.Printf("%s\n",res)
	t, _ := strconv.Atoi(string(res[(n-1)%b]))
	return t // 这个计算公式关键
}