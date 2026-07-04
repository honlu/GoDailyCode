package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var res int
	var left, right int
	subMap := make(map[byte]int)
	for left, right = 0, 0; right < len(s); {
		key := s[right]
		index, ok := subMap[key]
		if !ok {
			subMap[key] = right
			right++
			if res < len(subMap) {
				res = len(subMap)
			}
			continue
		}
		for ; left <= index; left++ {
			delete(subMap, s[left])
		}
		subMap[s[right]] = right
		right++

	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
