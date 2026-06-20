package p3

func lengthOfLongestSubstring(s string) int {
	count := map[byte]int{}
	left, res := 0, 0
	for right := 0; right < len(s); right++ {
		count[s[right]]++
		for count[s[right]] > 1 {
			count[s[left]]--
			left++
		}
		if length := right - left + 1; length > res {
			res = length
		}
	}
	return res
}
