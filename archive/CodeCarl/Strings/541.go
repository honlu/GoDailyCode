/*
541. 反转字符串2
2023-1-11
link: https://leetcode.cn/problems/reverse-string-ii/
question:
	给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，
	就反转这 2k 字符中的前 k 个字符。

	如果剩余字符少于 k 个，则将剩余字符全部反转。
	如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

answer:
*/
// 反转每2k中前k个字符串
func reverseStr(s string, k int) string {
	// 注意golang中string不能修改
	res := []byte(s)
	for i := 0; i < len(s); {
		if len(s)-i > k {
			reverseString(res[i : i+k])
		} else {
			reverseString(res[i:len(s)])
		}
		i += 2 * k
	}
	return string(res)
}

func reverseString(s []byte) {
	// 双指针
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}