package hot100

/*
3-最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/
func longestConsecutive(nums []int) int {
	// 首先去重
	numsMap := map[int]bool{}
	for _, num := range nums {
		numsMap[num] = true
	}
	var res int
	for num := range numsMap {
		if !numsMap[num-1] {
			cur := 1 // 当前num为开始的最大连续序列长度
			for numsMap[num+1] {
				num++
				cur++
			}
			if cur > res {
				res = cur
			}
		}
	}
	return res
}
