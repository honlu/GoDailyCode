/*
69. x的平方根
2022-11-24
link:https://leetcode.cn/problems/sqrtx/
question:
	给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
	由于返回类型是整数，结果只保留整数部分，小数部分将被舍去 。
answer:
	二分查找:注意细节和边界判断！
*/
func mySqrt(x int) int {
	low, high := 0, x
	var ans int // 添加一个保存真正mid的变量,即可以是最后mid的上一个值
	mid := low + (high-low)/2
	for low <= high { // 注意这里增加等于号的判断
		if mid*mid <= x {
			ans = mid
			low = mid + 1
		} else if mid*mid > x {
			high = mid - 1
		}
		mid = low + (high-low)/2
	}
	return ans
}
