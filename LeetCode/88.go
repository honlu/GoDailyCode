/*
88. 合并两个有序数组
2022-11-16
link: https://leetcode.cn/problems/merge-sorted-array/
question:
	给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，
	分别表示 nums1 和 nums2 中的元素数目。
	请你合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
answer:
	注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
	为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，
	后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
	做法一：将nums2中元素替代nums1尾部部分，然后排序!
	还有双指针的做法！再刷时拓展！

*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 替代后排序
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
}