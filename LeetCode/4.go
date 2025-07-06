package leetcode

import "fmt"

/*
4-寻找两个正序数组的中位数
超难、经典题。难不难，主要看做法.

思路：
1. 使用第三个数组保存两个有序数组的合并有序数组，然后再找中位值。
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

/*
思路2：递归做法
有很多边界问题，需要再review几遍
*/
