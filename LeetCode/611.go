/*
611. 有效三角形的个数
link： https://leetcode.cn/problems/valid-triangle-number/
question:
	给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。
answer:
	暴力枚举三个数，复杂度太高。要想办法降低复杂度。
	方法一：首先排序，固定最长的边，运用双指针扫描枚举。
*/
// 不能考虑枚举，太麻烦了。O(n^3)
func triangleNumber(nums []int) int {
	sort.Ints(nums) // 排序
	// 双指针
	n := len(nums)
	res := 0
	for i := n - 1; i >= 2; i-- {
		left, right := 0, i-1
		for left < right { // 这个for循环关键呀！
			if nums[left]+nums[right] > nums[i] {
				res += right - left
				right--
			} else {
				left++
			}
		}
	}
	return res
}