/*
16. 回文子串
2022-11-2
link: https://leetcode.cn/problems/palindromic-substrings/
question:
	给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
	具有不同开始位置或结束位置的子串，
	即使是由相同的字符组成，也会被视作不同的子串。
answer:
	动规五步曲：
	1、dp数组和下标含义：
		布尔类型的dp[i][j]：表示区间范围[i,j] （注意是左闭右闭）的子串是否是回文子串，
		如果是dp[i][j]为true，否则为false。
	2、递推公式
		当s[i]与s[j]不相等，则dp[i][j]一定是false。
		当s[i]与s[j]相等时，有如下三种情况:
		情况一：下标i 与 j相同，同一个字符例如a，当然是回文子串
		情况二：下标i 与 j相差为1，例如aa，也是回文子串
		情况三：下标i 与 j相差大于1的时候，这个区间是不是回文就看dp[i + 1][j - 1]是否为true。
	3、初始化：dp[i][j]初始化为false。
	4、遍历顺序：从下到上，从左到右遍历。可以保证dp[i + 1][j - 1]都是经过计算的。
		如果这矩阵是从上到下，从左到右遍历，那么会用到没有计算过的dp[i + 1][j - 1]，
		也就是根据不确定是不是回文的区间[i+1,j-1]，来判断了[i,j]是不是回文，那结果一定是不对的。
	5、举例推导
*/
func countSubstrings(s string) int {
	// dp数组和初始化
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := 0
	// 遍历
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 { // 情况一 和 情况二
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}