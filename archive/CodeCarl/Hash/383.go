/*
383. 赎金信
2023-1-10
link: https://leetcode.cn/problems/ransom-note/
question:
	给你两个字符串：ransomNote 和 magazine ，
	判断 ransomNote 能不能由 magazine 里面的字符构成。
	如果可以，返回 true ；否则返回 false 。
	magazine 中的每个字符只能在 ransomNote 中使用一次。

answerr:
	哈希，key为magazine中字符，value为magazine中字符出现的次数.
	注意：(还可以优化，代码未写)
		因为题目所只有小写字母，那可以采用空间换取时间的哈希策略，
		 用一个长度为26的数组还记录magazine里字母出现的次数.
*/
func canConstruct(ransomNote string, magazine string) bool {
	// map声明和初始化
	hash := map[rune]int{}
	for _, val := range magazine {
		hash[val]++
	}
	// 然后遍历ransomNote，查看是否符合要求
	for _, val := range ransomNote {
		hash[val]--
		if hash[val] < 0 {
			return false
		}
	}
	return true
}