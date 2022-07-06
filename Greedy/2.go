package greedy

import "sort"

/*
2、分发饼干
day:2022-6-17
link:https://leetcode.cn/problems/assign-cookies/
idea:
为了满足更多的小孩，就不要造成饼干尺寸的浪费。
大尺寸的饼干既可以满足胃口大的孩子也可以满足胃口小的孩子，那么就应该优先满足胃口大的。
这里的局部最优就是大饼干喂给胃口大的，充分利用饼干尺寸喂饱一个，全局最优就是喂饱尽可能多的小孩。
可以尝试使用贪心策略，先将饼干数组和小孩数组排序。
然后从后向前遍历小孩数组，用大饼干优先满足胃口大的，并统计满足小孩数量。
*/

// 贪心：先局部最有，进而推到全局最优
func findContentChildren(g []int, s []int) int {
	sort.Ints(g) // 排序孩子数组
	sort.Ints(s) // 排序饼干数组
	// 从小到大
	res := 0
	for i := 0; res < len(g) && i < len(s); i++ {
		if s[i] >= g[res] { // //如果饼干的大小大于或等于孩子的为空则给与，否则不给予，继续寻找选一个饼干是否符合
			res++
		}
	}
	return res
}
