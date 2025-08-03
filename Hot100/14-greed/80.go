package greed

/*
80-763:划分字母区间
https://leetcode.cn/problems/partition-labels/?envType=study-plan-v2&envId=top-100-liked
*/

/*
自己的解法思路：模拟（Simulation）+ 计数聚合判断 + 贪心思想（局部字符聚合贪心、聚合计数驱动的区间贪心）
「非最优解，后续可优化」
*/
func partitionLabels(s string) []int {
	// 记录每个字符在字符串中出现的总次数
	byteCount := map[byte]int{}
	for i := 0; i < len(s); i++ {
		byteCount[s[i]]++
	}

	var res []int
	curByte := map[byte]bool{} // 当前片段包含的字符集合
	start := 0                 // 当前片段的起始位置

	for i := 0; i < len(s); i++ {
		curByte[s[i]] = true

		// 如果当前片段长度恰好等于其中所有字符的总出现次数
		// 说明这些字符都只出现在这一段里，可以切割
		if i-start+1 == count(byteCount, curByte) { // 隐式贪心
			res = append(res, i+1-start)
			start = i + 1
			curByte = map[byte]bool{} // 重置片段字符集合
		}
	}
	return res
}

// 统计当前片段中所有字符的出现总次数
func count(byteCount map[byte]int, curByte map[byte]bool) int {
	sum := 0
	for ch := range curByte {
		sum += byteCount[ch]
	}
	return sum
}
