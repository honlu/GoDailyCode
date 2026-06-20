package dynamicprogramming

/*
90-32.最长有效括号
https://leetcode.cn/problems/longest-valid-parentheses?envType=study-plan-v2&envId=top-100-liked
*/
/*
定义：
dp[i] = 以 i 结尾的最长有效括号长度。

状态转移：
如果 s[i] == ')'：
如果前一个是 '('：dp[i] = dp[i-2] + 2
如果前一个是 ')'，并且在 i - dp[i-1] - 1 位置是 '('：
dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]

复杂度：O(n)，空间 O(n)。
*/
func longestValidParentheses(s string) int {
	// 注意只有两个符号
	n := len(s)
	dp := make([]int, n)
	var res int
	for i := 1; i < n; i++ {
		item := s[i]
		if item == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if s[i-1] == ')' {
				pre := i - dp[i-1] - 1
				if pre >= 0 && s[pre] == '(' {
					dp[i] = dp[i-1] + 2
					if pre-1 >= 0 {
						dp[i] += dp[pre-1]
					}
				}
			}
			if res < dp[i] {
				res = dp[i]
			}
		}
	}
	return res
}
