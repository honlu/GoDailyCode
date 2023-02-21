/*
17
738. 单调递增地数字
2022-10-13
link: 738-https://leetcode.cn/problems/monotone-increasing-digits/
question:

	给定一个非负整数 N，找出小于或等于 N 的最大的整数，
	同时这个整数需要满足其各个位数上的数字是单调递增。
	（当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，
	我们称这个整数是单调递增的。）

answer:

	局部最优：遇到strNum[i - 1] > strNum[i]的情况，让strNum[i - 1]--，
	然后strNum[i]给为9，可以保证这两位变成最大单调递增整数。
	全局最优：得到小于等于N的最大单调递增的整数。

	但这里局部最优推出全局最优，还需要其他条件，
	即后向遍历顺序，标记从哪一位开始统一改成9。
	【难点：这个标记后面都要改为9，是一定要的】
*/
package main

import (
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
	t := strconv.Itoa(n)
	s := []byte(t)
	// flag 用来标记赋值9从哪里开始
	flag := len(s)
	// 后向遍历
	for i := len(s) - 1; i > 0; i-- {
		if s[i-1] > s[i] {
			flag = i
			s[i-1]-- // 注意string是不能更改，所以要转为[]byte数组
		}
	}
	for i := flag; i < len(s); i++ {
		s[i] = '9'
	}
	res, _ := strconv.Atoi(string(s)) // 这个返回有两个值！
	return res
}
