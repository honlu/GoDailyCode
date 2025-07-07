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

// func longestPalindrome(s string) string {
// 	// 用于保存结果
// 	res := s[0:1]
// 	// dp数组创建和初始化
// 	dp := make([][]bool, len(s))
// 	for i := 0; i < len(s); i++ {
// 		dp[i] = make([]bool, len(s))
// 		dp[i][i] = true // 长度为1的子串初始化,其他的默认为false
// 	}
// 	// 遍历
// 	for i := 2; i <= len(s); i++ { // 固定子串的长度
// 		for j := 0; j < len(s)-i+1; j++ { //
// 			if s[j] != s[j+i-1] { // 如果首尾不同，则不可能为回文. j是子串的首部，j+i-1是子串的尾部
// 				continue
// 			} else if i < 3 {
// 				dp[j][j+i-1] = true // 即长度为2的判断
// 			} else { // 即s[j] = s[j+i-1]，首尾相同
// 				dp[j][j+i-1] = dp[j+1][j+i-2] // 状态转移
// 			}
// 			if dp[j][j+i-1] && i > len(res) { // 记录最大值
// 				res = s[j : j+i]
// 			}
// 		}
// 	}
// 	return res

// }

/*
做法思路：
1. 马拉车算法，类似KMP，做到O(n); 比较偏，只能做此题。
2. 双指针+哈希算法，有点难。不太适合最开始讲，o(log(n))
3. 比较简单做法：双指针解法
4. 两层for循环（O(n^3)）「下面解法」
*/
// func longestPalindrome(s string) string {
// 	var res string
// 	var max int
// 	// 暴力
// 	for i, j := 0, 0; i < len(s); i++ {
// 		for j := i; j < len(s); j++ {
// 			if palindrome(s[i:j+1]) && (j+1-i > max) {
// 				max = j + 1 - i
// 				res = string(s[i : j+1])
// 			}
// 		}
// 	}
// 	return res
// }

// func palindrome(item []byte) bool {
// 	if len(item) == 0 {
// 		return false
// 	}
// 	for i, j := 0, len(item)-1; i < j; {
// 		if item[i] != item[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }

/*
思路：双指针法, 减少一层回文串的判断，降低时间复杂度
假设从回文串的中心点，然后有左右两个指针、分别往左右走.
O(n^2)
*/
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		// 假设回文串为奇数串
		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if len(res) < r-1-l {
			res = string(s[l+1 : r])
		}
		// 假设回文串为偶数串
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if len(res) < r-1-l {
			res = string(s[l+1 : r])
		}
	}
	return res
}