/*
125. 验证回文串
2022-11-7
link: https://leetcode.cn/problems/valid-palindrome
question:
	如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，
	短语正着读和反着读都一样。则可以认为该短语是一个回文串。
	字母和数字都属于字母数字字符。
	给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
answer:
	首先进行字符串的处理，然后进行回文串的判断。
*/
func isPalindrome(s string) bool {
	// t保存字母和数字
	var temp []byte
	for i, _ := range s {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			temp = append(temp, s[i])
		}
	}
	t := strings.ToLower(string(temp)) // 转为小写
	// 回文串判断
	for i := 0; i < len(t)/2; i++ {
		if t[i] != t[len(t)-1-i] {
			return false
		}
	}
	return true
} 