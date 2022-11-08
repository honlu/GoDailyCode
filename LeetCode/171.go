/*
171. Excel表列序号
2022-11-8
link:https://leetcode.cn/problems/excel-sheet-column-number
question:
	给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。
	返回 该列名称对应的列序号 。
answer:
	模拟计算，遍历+计算
*/

func titleToNumber(columnTitle string) int {
	// 哈希表存储，也可以不存储，每次遍历时计算
	var hash map[byte]int = map[byte]int{'A': 1, 'B': 2, 'C': 3, 'D': 4,
		'E': 5, 'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10, 'K': 11, 'L': 12,
		'M': 13, 'N': 14, 'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20,
		'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26}
	// 从后往前模拟计算
	res := 0
	flag := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += flag * hash[columnTitle[i]]
		flag *= 26
	}
	return res
}