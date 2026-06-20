/*
169. 多数元素
2022-12-6
link: https://leetcode.cn/problems/majority-element/
question:
	给定一个大小为 n 的数组 nums ，返回其中的多数元素。
	多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
answer:
	你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
func majorityElement(nums []int) int {
	// 排序，在数组中间的数就是
	sort.Int(nums)
	return nums[len(nums)/2]
}