package hot100

import "fmt"

/*
9-找到字符串中所有字母异位词
思路：使用滑动窗口，维护一个窗口，窗口内的字符个数和p的字符个数相同
*/
func findAnagrams(s string, p string) []int {
	var ans []int
	if len(s) < len(p) {
		return ans
	}
	// 异位词的理解：两个字符串所含字符数目一样
	sCount := make(map[byte]int)
	pCount := make(map[byte]int)
	left, right := 0, len(p)
	// 统计滑动窗口中字符数
	for i := 0; i < right; i++ {
		sCount[s[i]]++
		pCount[p[i]]++
	}
	if equal(sCount, pCount) {
		ans = append(ans, left)
	}
	for right < len(s) {
		sCount[s[right]]++
		sCount[s[left]]--
		left++
		if left == 6 {
			fmt.Println(sCount, pCount)
		}
		if equal(sCount, pCount) {
			ans = append(ans, left)
		}
		right++
	}
	return ans
}

// 看两个滑动串口是否符合异位词
func equal(sCount, pCount map[byte]int) bool {
	// 注意sCount,pCount的元素个数可能不同，因为sCount中可能存在key真实存在，但value为0
	for key, value := range sCount {
		if value != pCount[key] {
			return false
		}
	}
	return true
}
