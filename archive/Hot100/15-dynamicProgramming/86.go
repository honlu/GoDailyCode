package dynamicprogramming

/*
86-139.单词拆分
https://leetcode.cn/problems/word-break/?envType=study-plan-v2&envId=top-100-liked
*/
/*
1. 暴力解法：
最原始的想法：
s 从开头截一部分，如果这部分在词典里 → 剩下的继续拆
一直拆到末尾，全部匹配就成功。
这其实就是 递归/DFS。
问题：会超时（因为有很多重复拆分），无法通过所有case。
*/
// func wordBreak(s string, wordDict []string) bool {
// 	// 暴力接试试，首先将字典转为map
// 	wordMap := make(map[string]bool)
// 	for i := range wordDict {
// 		wordMap[wordDict[i]] = true
// 	}
// 	var dfs func(item string) bool
// 	dfs = func(item string) bool {
// 		if len(item) == 0 {
// 			return true
// 		}
// 		// 拆成各个片段判断
// 		for i := 1; i <= len(item); i++ { // 暴力需要注意<=有等于
// 			if wordMap[item[:i]] && dfs(item[i:]) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	return dfs(s)
// }

/*
2. 递归+记忆性。可以通过所有case
*/
// func wordBreak(s string, wordDict []string) bool {
// 	// 把字典转为map，方便O(1)判断
// 	wordMap := make(map[string]bool)
// 	for _, w := range wordDict {
// 		wordMap[w] = true
// 	}

// 	memo := make(map[string]bool) // 记忆化，存储子串结果

// 	var dfs func(item string) bool
// 	dfs = func(item string) bool {
// 		if len(item) == 0 {
// 			return true
// 		}
// 		if v, ok := memo[item]; ok { // 如果之前算过，直接返回
// 			return v
// 		}
// 		for i := 1; i <= len(item); i++ {
// 			if wordMap[item[:i]] && dfs(item[i:]) {
// 				memo[item] = true
// 				return true
// 			}
// 		}
// 		memo[item] = false
// 		return false
// 	}

// 	return dfs(s)
// }

/*
3. 动态规划（DP）
我们用 一维 DP：
定义：dp[i] 表示前 i 个字符能否被拆分。
转移方程：
dp[i] = true if 存在 j < i，使得 dp[j] == true 且 s[j:i] 在字典中
初始值：dp[0] = true （空串合法）。
答案：dp[n]。
*/
// 复杂度 O(n^2)，比 DFS 更直观
func wordBreakDP(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true // 空串合法

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
