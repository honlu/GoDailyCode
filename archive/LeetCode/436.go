package leetcode

/*
436-寻找右区间
https://leetcode.cn/problems/find-right-interval/description/
*/
/*
思路：暴力for两层解决。
需要注意两点：1. 结果索引可能是本身 2. 寻找找到最小的右区间索引
*/
func findRightInterval(intervals [][]int) []int {
	m := len(intervals)
	res := make([]int, m)
	for i := 0; i < m; i++ {
		res[i] = -1
	}
	for i := 0; i < m; i++ {
		endI := intervals[i][1]
		for j := 0; j < m; j++ {
			startJ := intervals[j][0]
			if startJ >= endI {
				if res[i] == -1 || (res[i] != -1 && startJ < intervals[res[i]][0]) {
					res[i] = j
				}
			}
		}
	}
	return res
}
