package lcr183

import "math"

/*
题目：望远镜中最高的海拔
- https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 科技馆内有一台虚拟观景望远镜，它可以用来观测特定纬度地区的地形情况。该纬度的海拔数据记于数组 heights ，其中 heights[i] 表示对应位置的海拔高度。请找出并返回望远镜视野范围 limit 内，可以观测到的最高海拔值。

- 标签：队列，滑动窗口

代码还可以优化，有时间再搞。
*/

func maxAltitude(heights []int, limit int) []int {
	var stack []int
	var res []int
	if limit == 1 || limit == 0 { // 特殊情况
		return heights
	}
	if len(heights) == limit { // 特殊情况2
		res = append(res, max(heights))
		return res
	}
	for i := 0; i < len(heights); i++ {
		if i+1 < limit {
			stack = append(stack, heights[i])
		} else if i+1 == limit { // 边界问题
			stack = append(stack, heights[i]) // 进
			res = append(res, max(stack))     // 最大值
		} else {
			stack = stack[1:]                 // 出
			stack = append(stack, heights[i]) // 进
			res = append(res, max(stack))     // 最大值
		}
	}

	return res
}

func max(stack []int) int {
	var res int = math.MinInt32
	for _, item := range stack {
		if res < item {
			res = item
		}
	}
	return res
}
