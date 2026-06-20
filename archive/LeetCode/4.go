package leetcode

/*
4-寻找两个正序数组的中位数
超难、经典题。难不难，主要看做法.

思路：
1. 使用第三个数组保存两个有序数组的合并有序数组，然后再找中位值。
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	// 使用第三个数组保存有序数组，然后取中位
// 	count := len(nums1) + len(nums2)
// 	var nums3 []int
// 	l1, l2 := 0, 0
// 	for l1 < len(nums1) && l2 < len(nums2) {
// 		if nums1[l1] < nums2[l2] {
// 			nums3 = append(nums3, nums1[l1])
// 			l1++
// 		} else {
// 			nums3 = append(nums3, nums2[l2])
// 			l2++
// 		}
// 	}
// 	if l1 < len(nums1) {
// 		nums3 = append(nums3, append([]int{}, nums1[l1:]...)...)
// 	}
// 	if l2 < len(nums2) {
// 		nums3 = append(nums3, append([]int{}, nums2[l2:]...)...)
// 	}
// 	fmt.Println(nums3)
// 	if (count)%2 == 0 {
// 		return float64(nums3[count/2-1]*1.0+nums3[count/2]*1.0) / 2.0
// 	}
// 	return float64(nums3[count/2])
// }

/*
思路2：递归做法，解法为寻找两个有序数组中第k个大小的值
有很多边界问题，需要再review几遍.
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 递归寻找
	count := len(nums1) + len(nums2)
	var find func(nums1, nums2 []int, i, j, k int) int // 寻找两个从i,j开始的两个有序数组中，第k个大小的值
	find = func(nums1, nums2 []int, i, j, k int) int {
		if len(nums1)-i > len(nums2)-j { // 保证第一个是短串，第二是长串
			return find(nums2, nums1, j, i, k)
		}
		if k == 1 { // 边界
			if len(nums1) == i {
				return nums2[j]
			} else {
				return min(nums1[i], nums2[j])
			}
		}
		if len(nums1) == i { // 边界
			return nums2[j+k-1]
		}
		si, sj := min(len(nums1), i+k/2), j+k-k/2
		if nums1[si-1] > nums2[sj-1] {
			return find(nums1, nums2, i, sj, k-(sj-j))
		} else {
			return find(nums1, nums2, si, j, k-(si-i))
		}

	}
	if count%2 == 0 {
		left := find(nums1, nums2, 0, 0, count/2)
		right := find(nums1, nums2, 0, 0, count/2+1)
		return (float64(left) + float64(right)) / 2.0
	}
	return float64(find(nums1, nums2, 0, 0, count/2+1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
