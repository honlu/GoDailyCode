/*
5. 最长回文子串
2022-11-14
link: https://leetcode.cn/problems/longest-palindromic-substring/
question:
	给你一个字符串 s，找到 s 中最长的回文子串。
answer:
	动态规划五步曲：
	1、dp数组和下标含义：
		dp[i][j]表示从s[i...j]的字符串是否为回文，true是，false不是
	2、递推公式
		dp[i][j] = dp[i+1][j-1] && (s[i]==s[j])
	3、初始化
		子串为1的都为true，即dp[i][i]=true
	4、遍历顺序[关键]
		对角填充，通过固定子串的长度，然后一一遍历起点。
	5、举例推导
*/

func longestPalindrome(s string) string {
	// 用于保存结果
	res := s[0:1]
	// dp数组创建和初始化
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 长度为1的子串初始化,其他的默认为false
	}
	// 遍历
	for i := 2; i <= len(s); i++ { // 固定子串的长度
		for j := 0; j < len(s)-i+1; j++ { //
			if s[j] != s[j+i-1] { // 如果首尾不同，则不可能为回文. j是子串的首部，j+i-1是子串的尾部
				continue
			} else if i < 3 {
				dp[j][j+i-1] = true // 即长度为2的判断
			} else { // 即s[j] = s[j+i-1]，首尾相同
				dp[j][j+i-1] = dp[j+1][j+i-2] // 状态转移
			}
			if dp[j][j+i-1] && i > len(res) { // 记录最大值
				res = s[j : j+i]
			}
		}
	}
	return res

}
