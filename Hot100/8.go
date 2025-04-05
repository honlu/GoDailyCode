package hot100

/*
8-无重复字符的最长子串
滑动窗口
*/
func lengthOfLongestSubstring(s string) int {
	var ans int
	windows := make(map[byte]int) // 窗口内是否存在某个字符的个数map
	// 滑动窗口如何使用？注意使用注意事项
	var left, right int
	for right < len(s) { // 注意这里使用的窗口，left, right分别是左闭右开型
		cur := s[right] // 当前要加入到窗口的字符
		windows[cur]++
		right++
		for windows[cur] > 1 { // 若当前元素在原窗口内已经存在，要移除旧窗口内当前元素和左边的所有元素
			remove := s[left]
			windows[remove]--
			left++
		}
		ans = max(ans, right-left)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
