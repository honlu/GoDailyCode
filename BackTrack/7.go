/*
7. 分割回文串
day: 2022-10-7
link: 131-https://leetcode.cn/problems/palindrome-partitioning/
question:
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
	返回 s 所有可能的分割方案。
answer:
	其实切割问题类似组合问题.也抽象成一个树形结构来解决。
	递归用来纵向遍历，for循环用来横向遍历，切割到字符串的结尾位置，说明找到了一个切割方法。
*/
func partition(s string) [][]string {
	var track []string // 切割字符串的集合
	var res [][]string // 结果集合
	backTrack(0, s, track, &res)
	return res
}

// 递归函数-回溯
func backTrack(startIndex int, s string, track []string, res *[][]string) {
	if startIndex == len(s) { // 到达字符串末尾
		temp := make([]string, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}
	// for 循环
	for i := startIndex; i < len(s); i++ {
		// 处理。首先判断startIndex和i切割的空间是否为回文，若为回文，则加入到track,否则继续后移，找到回文区间
		if isPartition(s, startIndex, i) {
			track = append(track, s[startIndex:i+1])
		} else {
			continue
		}
		// 递归
		backTrack(i+1, s, track, res) // 注意这里res
		// 回溯
		track = track[:len(track)-1]
	}
}

// 判断是否是回文
func isPartition(s string, startIndex, end int) bool {
	left, right := startIndex, end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}