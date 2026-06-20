package multidimensionaldynamicprogramming

/*
93-5. 最长回文子串
https://leetcode.cn/problems/longest-palindromic-substring/?envType=study-plan-v2&envId=top-100-liked
*/
/*
动态规划：
dp[i][j]代表s[i:j]是否为回文子串
转移方式：
dp[i][j] = s[i]==s[j] &&(j-i<3 || dp[i+1][j-1])
（长度 <= 2 时只需两端相等即可，否则看内层是否回文）
初始化：
dp[i][i]=true
注意：
边界，以及其他
*/
func longestPalindrome(s string) string {
	m := len(s)
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, m)
		dp[i][i] = true
	}
	start, maxL := 0, 1       // 最长回文子串的起点和长度
	for l := 2; l <= m; l++ { // 回文串长度,注意<=m
		for i := 0; i <= m-l; i++ { // 起点,注意<=m-l
			j := i + l - 1
			if s[i] == s[j] {
				if l <= 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1] // 这样合理吗？合理，因为首先遍历了短长度的回文串
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] && l > maxL {
				start, maxL = i, l
			}
		}
	}
	return s[start : start+maxL]
}
