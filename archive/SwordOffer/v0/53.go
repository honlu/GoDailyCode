/*
剑指offer53-1: 在排序数组中查找数字I
2023-4-15 by lu
link: https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
question:
	统计一个数字在排序数组中出现的次数。
answer:
	主要考察查找！
	可以使用哈希等，golang里面有针对排序的sort.SearchInts可以使用
*/
func search(nums []int, target int) int {
	index := sort.SearchInts(nums, target)
	res := 0
	if index == len(nums) {
		return res
	}
	for index < len(nums) {
		if nums[index] == target {
			res++
			index++
		} else {
			break
		}
	}
	return res
}