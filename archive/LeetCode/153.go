/*
153. 寻找旋转排序数组中的最小值
2022-12-16
link: https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
question:
	已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。
answer:
	二分查找找到最小元素的位置，即左大又小，或者在边界。
*/
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid = low + (high-low)/2
		if nums[mid] < nums[high] { // 判断哪边是有序的？
			high = mid // mid-1不行,因为此时mid的左边（包含）是不完全有序的那一半，mid有可能直接是最小值，所以要取mid
		} else {
			low = mid + 1 // #将范围缩小到无序的那一半，因为答案就在那一半。之所以要+1，
			// 是因为mid肯定不是最小的那个，至少nums[right]比nums[mid]更小
		}
	}
	return nums[low]
}