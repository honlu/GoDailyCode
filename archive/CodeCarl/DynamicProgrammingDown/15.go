/*
15. 编辑举例
2022-11-2
link:https://leetcode.cn/problems/edit-distance/
question:
	给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
answer:
	编辑距离是用动规来解决的经典题目，
	这道题目看上去好像很复杂，但用动规可以很巧妙的算出最少编辑距离。
	五步曲：
	1、dp数组和下标含义：[关键]
		dp[i][j] 表示以下标i-1为结尾的字符串word1，和以下标j-1为结尾的字符串word2，最近编辑距离为dp[i][j]。
		为啥要表示下标i-1为结尾的字符串呢，为啥不表示下标i为结尾的字符串呢？
		用i来表示也可以！ 但我统一以下标i-1为结尾的字符串，在下面的递归公式中会容易理解一点。
	2、递推公式[关键]
		word1[i - 1] 与 word2[j - 1]相等了，那么就不用编辑了，
		以下标i-2为结尾的字符串word1和以下标j-2为结尾的字符串word2的最近编辑距离dp[i - 1][j - 1]就是 dp[i][j]了。
		当 if (word1[i - 1] != word2[j - 1]) 时取最小的，即：dp[i][j] = min({dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1]}) + 1;
	3、初始化
		dp[i][0]就应该是i，对word1里的元素全部做删除操作，即：dp[i][0] = i;
		同理dp[0][j] = j;
	4、遍历顺序:从左到右从上到下去遍历。
	5、举例推导
*/

func minDistance(word1 string, word2 string) int {
	// dp数组和初始化
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
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
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1 // 注意这里和“两个字符串删减操作”题代码的唯一区别
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