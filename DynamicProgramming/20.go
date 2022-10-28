/*
20. 单词拆分
2022-10-28
link: https://leetcode.cn/problems/word-break/
question:
	给你一个字符串 s 和一个字符串列表 wordDict 作为字典。
	请你判断是否可以利用字典中出现的单词拼接出 s 。
answer:
	注意：拆分时可以重复使用字典中的单词。假设字典中没有重复的单词。
	背包问题角度来看：
		将单词看出物品，字符串s看成背包。单词能否组成字符串s，就是问物品是否可以将背包装满。
		单词可以重复使用，完全背包！
	动态规划五部曲：
	1、确定dp数组和下标含义：dp[i],字符串长度为i的话，dp[i]为true,表示可以拆分为一个或多个字典中出现的单词。
	2、递推公式：
	如果dp[j]=true,j <i,且[j,i]这个区间的字串出现在字典里，那么dp[i]一定是true
	所以递推公式： if [j,i]区间子串在字典中 && dp[j]== true: dp[i]=true
	3、dp数组初始化：
	dp[0]为true, 其他下标为false.
	4、遍历顺序：本题最终问是否都出现过，所以对出现单词集合里的元素是组合还是排列，并不在意！
	两种遍历顺序都可以。但有个特殊性，因为要求子串，最好是遍历背包放在外循环，将遍历物品放在内循环。
	5、举例推到dp
*/
func wordBreak(s string, wordDict []string) bool {
	// 哈希表，方便判断子串是否在字典中
	hash := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		hash[word] = true
	}
	// 创建dp数组和初始化
	dp := make([]bool, len(s)+1)
	dp[0] = true
	// 遍历
	for i := 1; i <= len(s); i++ { // 遍历背包
		for j := 0; j < i; j++ { // 遍历物品
			subStr := s[j:i]
			if hash[subStr] == true && dp[j] == true {
				dp[i] = true
				break // 这里可以跳出内循环，不写也可以。
			}
		}
	}
	return dp[len(s)] // 注意没有-1
}