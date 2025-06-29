package binarysearch

/*
34-在排序数组中查找元素的第一个和最后一个位置
https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked

难点：
- 看到先想要：「二分先找到某个元素控制，然后往前或往后遍历」，结果：测试运行超时.
- 暴力解法：先找区间的开始节点。
- to-do：有时间重新优化二分的解法。
*/
func searchRange(nums []int, target int) []int {
	// 暴力，先找左区间，然后到右区间。
	left, right := -1, -1
	var res = []int{left, right}
	for i, val := range nums {
		if val == target {
			if left == -1 {
				left = i
				right = i
			} else {
				right = right + 1
			}
		}
	}
	res[0], res[1] = left, right
	return res
}
