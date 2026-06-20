/*
14
435. 无重叠部分
2022-10-13
2023-2-19 updated by lu
link: https://leetcode.cn/problems/non-overlapping-intervals/
question:
	给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。
	返回 需要移除区间的最小数量，使剩余区间互不重叠
answer:
	注意: 可以认为区间的终点总是大于它的起点。
	区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
	思路：按照左边界排序，从左向右记录交叉区间的个数。

*/
func eraseOverlapIntervals(intervals [][]int) int {
	// 排序(左边界)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	if len(intervals) == 0 {
		return 0
	}
	// 遍历比较
	res := 0
	for i := 1; i < len(intervals); i++ { // 里面比较是关键和技巧
		if intervals[i][0] < intervals[i-1][1] { // 有交叉区间
			res++
			if intervals[i-1][1] < intervals[i][1] { // 更改
				intervals[i][1] = intervals[i-1][1]
			}
		}
	}
	return res
}

// 保存不重叠的区间
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 排序-左边界排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	var temp [][]int // 保存不重叠的区间
	temp = append(temp, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if temp[len(temp)-1][1] <= intervals[i][0] {
			temp = append(temp, intervals[i])
		} else if temp[len(temp)-1][1] > intervals[i][0] {
			if temp[len(temp)-1][1] > intervals[i][1] { // 这里是关键
				temp = temp[:len(temp)-1]
				temp = append(temp, intervals[i])
			}
		}
	}
	return len(intervals) - len(temp)
}