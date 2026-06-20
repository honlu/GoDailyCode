package main

import (
	"sort"
)

func checkDynasty(places []int) bool {
	// 查看有替代性数字的数组是否连续
	// 排序+最大小限制
	sort.Ints(places)
	var unknow int
	for i := 1; i < len(places); i++ {
		if places[i-1] == 0 {
			unknow++
			continue
		} else if places[i] == places[i-1] {
			return false
		}
	}
	return (places[4] - places[unknow]) < 5
}
