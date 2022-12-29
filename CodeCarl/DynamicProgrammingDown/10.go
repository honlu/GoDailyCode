/*
10. 不相交的线
2022-11-2
link:
question:
	我们在两条独立的水平线上按给定的顺序写下 A 和 B 中的整数。
	现在，我们可以绘制一些连接两个数字 A[i] 和 B[j] 的直线，
	只要 A[i] == B[j]，且我们绘制的直线不与任何其他连线（非水平线）相交。
	以这种方法绘制线条，并返回我们可以绘制的最大连线数。
answer:
	绘制一些连接两个数字 A[i] 和 B[j] 的直线，只要 A[i] == B[j]，且直线不能相交！
	直线不能相交，这就是说明在字符串A中找到一个与字符串B相同的子序列，
	且这个子序列不能改变相对顺序，只要相对顺序不改变，连接相同数字的直线就不会相交。
	其实也就是说A和B的最长公共子序列是[1,4]，长度为2。 这个公共子序列指的是相对顺序不变。
	（即数字4在字符串A中数字1的后面，那么数字4也应该在字符串B数字1的后面）
	所以：本题说是求绘制的最大连线数，其实就是求两个字符串的最长公共子序列的长度！
	代码和“最长公共子序列”一样，除了字符串名不一样而已。
*/
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	// dp数组和初始化
	dp := make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	// 遍历
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}