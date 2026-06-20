/*
15
763.划分字母区间
2022-10-13
2023-2-20 updated by lu
link: https://leetcode.cn/problems/partition-labels/
question:
	字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，
	同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
answer:
	思路：
	统计每一个字符最后出现的位置.
	从头遍历字符，并更新字符的最远出现下标，【这个思路非常重要，不容易想到】
	如果找到字符最远出现位置下标和当前下标相等了，则找到了分割点
难点：如何把同一个字母的都圈在同一个区间里呢？
如果没有接触过这种题目的话，还挺有难度的。
*/

func partitionLabels(s string) []int {
	hash := make(map[byte]int, 27)
	for i := 0; i < len(s); i++ { // // 统计每一个字符最后出现的位置
		hash[s[i]] = i
	}
	var res []int
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if hash[s[i]] > right {
			right = hash[s[i]]
		}
		if i == right { // 找到字符最远出现位置下标和当前下标相等
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

// 代码优化，不适用map，仅仅数组保存，节省内存
func partitionLabels(s string) []int {
	lastPos := [26]int{} // 统计每个字母出现的最远地址
	for i, c := range s {
		lastPos[c-'a'] = i
	}

	start, right := 0, 0 // 一个片段的起始和结束地址
	var res []int
	for i, c := range s {
		if lastPos[c-'a'] > right {
			right = lastPos[c-'a']
		}
		if i == right {
			res = append(res, right-start+1)
			start = right + 1
		}
	}
	return res
}