/*
16、合并区间
2022-10-13
link:56-https://leetcode.cn/problems/merge-intervals/
question:
	给出一个区间的集合，请合并所有重叠的区间。
answer:
	和13.go\14.go很类似。代码稍微改一改就可以了。
*/
func merge(intervals [][]int) [][]int {
	// 排序(左边界)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 遍历比较
	for i := 0; i < len(intervals)-1; i++ { // 里面比较是关键和技巧
		// 比较前后区间，如果重合先更改前一个区间地尾部
		if intervals[i][1] >= intervals[i+1][0] { // 有交叉区间。注意大于等于号，等于号不能少
			if intervals[i][1] < intervals[i+1][1] { // 更改
				intervals[i][1] = intervals[i+1][1]
			}
			// 然后，删除后一个区间(即i+1)
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			// 重新比较i和i+1
			i--
		}
	}
	return intervals
}