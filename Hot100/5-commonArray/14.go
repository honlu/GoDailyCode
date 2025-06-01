package hot100

import "sort"

/*
14-合并区间和
*/
func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 {
		return res
	}
	// 自定义数组排序算法+遍历
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] { // 注意合并
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
