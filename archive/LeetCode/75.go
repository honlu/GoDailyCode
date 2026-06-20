/*
75. 颜色分类
link: https://leetcode.cn/problems/sort-colors/
question:
	给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，
	使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
	我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
	必须在不使用库内置的 sort 函数的情况下解决这个问题。
answer:
	双指针思想，但具体实现细节可以有不同
*/
// 利用partition划分区域思想实现
func sortColors(nums []int) {
	left, right := -1, len(nums)
	// 选定1作为基准值，将其化成三个区域，大于1的区域，小于1的区域，等于1的区域
	for i := 0; i < right; {
		if nums[i] > 1 { // 右边放大于1的区域
			right--
			nums[i], nums[right] = nums[right], nums[i]
			// 注意后面不要i++,因为不确定交换后i索引的值，需要继续判断
		} else if nums[i] < 1 { // 左边放小于1的区域
			left++
			nums[i], nums[left] = nums[left], nums[i]
			i++
		} else { // 中间是1的区域
			i++
		}
	}
}

// 利用循环，先抽选2，再排1的思想
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	p0 := 0
	for left <= right {
		for left <= right && nums[left] == 2 { // 先将2排号
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
		if nums[left] == 0 {
			nums[left], nums[p0] = nums[p0], nums[left]
			p0++
		}
		left++
	}
}
 