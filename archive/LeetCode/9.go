package leetcode

/*
9. 回文数
2023-3-16
link: https://leetcode.cn/problems/palindrome-number/
question:
	给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
	回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
	例如，121 是回文，而 123 不是。

answer:
	方法一：整数转为字符串来解决！二分
	方法二：将数字反转，然后对比两个数字是否相等
*/
// func isPalindrome(x int) bool {
// 	// 转为字符串来解决
// 	s := strconv.Itoa(x)
// 	n := len(s)
// 	for i := 0; i < n/2; i++ {
// 		if s[i] != s[n-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }
/*
转为字符串
*/
// func isPalindrome(x int) bool {
// 	// 字符串就可以
// 	s := strconv.Itoa(x)
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 	}
// 	return true
// }

// 方法二：数字反转
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	y := 0
	for temp > 0 {
		y = y*10 + temp%10
		temp /= 10
	}
	return y == x
}
