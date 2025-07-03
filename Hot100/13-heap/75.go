package heap

import "sort"

/*
75-前K个高频元素
https://leetcode.cn/problems/top-k-frequent-elements/description/?envType=study-plan-v2&envId=top-100-liked
*/

func topKFrequent(nums []int, k int) []int {
	// 先统计每个元素出现频次，在排序
	numMap := make(map[int]int)
	items := []int{}
	for i, _ := range nums {
		item := nums[i]
		if numMap[item] > 0 {
			numMap[item]++
		} else {
			numMap[item]++
			items = append(items, item)
		}
	}
	sort.Slice(items, func(a, b int) bool {
		return numMap[items[a]] > numMap[items[b]]
	})
	return items[:k]
}
