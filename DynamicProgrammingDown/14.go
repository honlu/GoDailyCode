/*
14. 两个字符串的删减操作
2022-11-2
link: https://leetcode.cn/problems/delete-operation-for-two-strings/
question:
	给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，
	每步可以删除任意一个字符串中的一个字符。
answer:
	方法一：
	本题和动态规划：115.不同的子序列 (opens new window)相比，其实就是两个字符串都可以删除了，
	情况虽说复杂一些，但整体思路是不变的。
	这次是两个字符串可以相互删了，这种题目用动态规划的思路来解.
	动态规划五步曲：
	1、dp数组和下标含义[关键，有点绕]
		dp[i][j]：以i-1为结尾的字符串word1，
		和以j-1位结尾的字符串word2，想要达到相等，所需要删除元素的最少次数。
	2、递推公式
		当word1[i - 1] 与 word2[j - 1]相同的时候，dp[i][j] = dp[i - 1][j - 1];
		当word1[i - 1] 与 word2[j - 1]不相同的时候：
		递推公式：dp[i][j] = min({dp[i - 1][j - 1] + 2, dp[i - 1][j] + 1, dp[i][j - 1] + 1});
		因为dp[i - 1][j - 1] + 1等于 dp[i - 1][j] 或 dp[i][j - 1]，
		所以递推公式可简化为：dp[i][j] = min(dp[i - 1][j] + 1, dp[i][j - 1] + 1);
	3、初始化
		dp[i][0]：word2为空字符串，以i-1为结尾的字符串word1要删除多少个元素，才能和word2相同呢，很明显dp[i][0] = i。
	4、遍历顺序
		从上到下，从左到右，这样保证dp[i][j]可以根据之前计算出来的数值进行计算。
	5、举例推导
*/

// 方法一
func minDistance(word1 string, word2 string) int {
	// dp数组和初始化
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i // 初始化，有点绕
	}
	for i := 0; i <= len(word2); i++ {
		dp[0][i] = i
	}
	// 遍历
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}