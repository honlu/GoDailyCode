/*
395. 至少有k个重复字符的最长子串
2023-4-17 by lu
link: https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/
question:
	给你一个字符串 s 和一个整数 k ，
	请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
answer:
	虽然中等，但感觉是困难题！参考题解解答。
	主要分治的思路不好想，以及代码细节实现很重要！
	https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/solution/shou-hua-tu-jie-tong-guo-fen-zhi-suo-xia-cnn1/
*/
func longestSubstring(s string, k int) int {
	// 分治
	return helper(0, len(s)-1, k, s)
}

// 需要好好理解细节
func helper(start, end, k int, s string) int {
	if end-start+1 < k {
		return 0
	}

	freq := make(map[byte]int, end-start+1) // 哈希
	for i := start; i <= end; i++ {         // 全局来看各个字符个数
		freq[s[i]]++
	}
	for end-start+1 >= k && freq[s[start]] < k { // 找出合法的子串
		start++
	}
	for end-start+1 >= k && freq[s[end]] < k {
		end--
	}

	if end-start+1 < k {
		return 0
	}

	for i := start; i <= end; i++ {
		if freq[s[i]] < k {
			return max(helper(start, i-1, k, s), helper(i+1, end, k, s))
		}
	}
	return end - start + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}