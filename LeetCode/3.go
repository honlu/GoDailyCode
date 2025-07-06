package leetcode

/*
3-无重复字符的最长子串
https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/

涉及：双指针算法、滑动窗口（双指针）
双指针：最多时间复杂度O(2n)即O(n)
*/
// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 1 {
// 		return 1
// 	}
// 	// 滑动窗口(双指针算法)，哈希。
// 	left, right := 0, 0
// 	valueIndex := make(map[byte]int)
// 	var max int
// 	// 看每个索引作为起点的最长子串
// 	for right < len(s) {
// 		if _, ok := valueIndex[s[right]]; ok {
// 			left++
// 			right = left
// 			valueIndex = make(map[byte]int)
// 		} else {
// 			if right-left+1 > max { // 注意这里+1
// 				fmt.Println(right, left)
// 				max = right - left + 1
// 			}
// 			valueIndex[s[right]] = right
// 			right++ // 注意这里才有+1
// 		}
// 	}
// 	return max
// }

/* 
优化代码逻辑
*/
func lengthOfLongestSubstring(s string) int {
    var res int
	valueMap := make(map[byte]int)
	for right, left := 0, 0; right < len(s); right++ {
		valueMap[s[right]]++
		for valueMap[s[right]] > 1 { 
			// 注意这里必须大于1，说明当前以right节点的子串有重复字符了
			// 然后移动左节点，去重
			valueMap[s[left]]--
			left++
		}
		if res < (right- left + 1) {
			res = right - left + 1
		}
	}
	return res
}
}
