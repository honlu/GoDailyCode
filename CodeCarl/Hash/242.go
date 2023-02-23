/*
242. 有效的字母异位词
2023-1-8
link: https://leetcode.cn/problems/valid-anagram/
question:
	给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
	注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

answer:
	简单方式先排序，遍历判断。
	进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
		关于unicode字符核心点在于「字符是离散未知的」，
		因此我们用哈希表维护对应字符的频次即可。
		同时读者需要注意Unicode 一个字符可能对应多个字节的问题，
		不同语言对于字符串读取处理的方式是不同的。

*/
// 排序+遍历
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	// 字符串排序
// 	r1 := []rune(s)
// 	r2 := []rune(t)
// 	sort.Slice(r1, func(i, j int) {
// 		return r1[i] < r1[j]
// 	})
// 	sort.Slice(r2, func(i, j int) {
// 		return r2[i] < r2[j]
// 	})
// 	// 遍历. 不遍历也可以：字符串直接判断 string(r1) == string(r2)
// 	for i := 0; i < len(s); i++ {
// 		if r1[i] != r2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 哈希表
	hash := map[rune]int{}
	for _, val := range s {
		hash[val]++
	}
	for _, val := range t {
		hash[val]--
		if hash[val] < 0 {
			return false
		}
	}
	return true

}