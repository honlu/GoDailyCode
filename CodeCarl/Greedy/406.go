/*
12
406.根据身高重建队列
2022-10-13
2023-2-17 update by lu
link: https://leetcode.cn/problems/queue-reconstruction-by-height/
question:
	假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。
answer:
	提醒：遇到两个维度权衡的时候，一定要先确定一个维度，再确定另一个维度。
	如果两个维度一起考虑一定会顾此失彼。

	本体思路：先按照身高从高到低排序，让高个子在前面。然后按照k为下标重新插入队列！

	局部最优：优先按身高高的people的k来插入。插入操作过后的people满足队列属性
全局最优：最后都做完插入操作，整个队列满足题目队列属性
*/
package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool { // 按照身高从大到小排序，按照位置从小到大
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return false
	})
	// 插入。不建议直接在people上操作，新建一个queue
	queue := make([][]int, 0)
	for i := 0; i < len(people); i++ {
		index := people[i][1]
		// 在index处插入
		queue = append(queue[:index], append([][]int{people[i]}, queue[index:]...)...)
	}
	return queue
}
