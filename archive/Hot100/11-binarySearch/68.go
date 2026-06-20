package binarysearch

import "fmt"

/*
68-4.寻找两个正序数组的中位数
https://leetcode.cn/problems/median-of-two-sorted-arrays/description/?envType=study-plan-v2&envId=top-100-liked
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 使用第三个数组保存有序数组，然后取中位
	count := len(nums1) + len(nums2)
	var nums3 []int
	l1, l2 := 0, 0
	for l1 < len(nums1) && l2 < len(nums2) {
		if nums1[l1] < nums2[l2] {
			nums3 = append(nums3, nums1[l1])
			l1++
		} else {
			nums3 = append(nums3, nums2[l2])
			l2++
		}
	}
	if l1 < len(nums1) {
		nums3 = append(nums3, append([]int{}, nums1[l1:]...)...)
	}
	if l2 < len(nums2) {
		nums3 = append(nums3, append([]int{}, nums2[l2:]...)...)
	}
	fmt.Println(nums3)
	if (count)%2 == 0 {
		return float64(nums3[count/2-1]*1.0+nums3[count/2]*1.0) / 2.0
	}
	return float64(nums3[count/2])
}
