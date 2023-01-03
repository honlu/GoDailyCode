/*
27. 移除元素
2023-1-3
link: https://leetcode.cn/problems/remove-element/
question:
	给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
	不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
	元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

answer:
	实际要求可以理解为是将val值的元素移动数组最后面或者被覆盖，不是val的值在数组前面。返回前面数组不是val的值的长度。

	暴力解法可以：两层for循环，一个for循环遍历数组元素，当发现val值时，第二个for循环更新数组,将数组后面的全部元素前向移动一位。

	双指针（快慢指针）： 通过一个快指针和慢指针在一个for循环下完成两个for循环的工作
		快指针：寻找新数组的元素 ，新数组就是不含有目标元素的数组
		慢指针：指向更新 新数组下标的位置
*/

// 暴力解法
func removeElement(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			for j := i + 1; j < size; j++ { // 后面元素全部向前移动一位
				nums[j-1] = nums[j]
			}
			i-- // 这里细节！因为下标i以后的数值都向前移动了一位，所以i也向前移动一位，为了保持一致。因为外层for循环最后还有i++
			size--
		}
	}
	return size
}

// 双指针: 下边定义成左右指针更好理解
// 快指针：寻找新数组的元素 ，新数组就是不含有目标元素的数组
// 慢指针：指向更新 新数组下标的位置
func removeElement(nums []int, val int) int {
	low, fast := 0, len(nums)-1 // 注意这里使用第1种区间定义[low, fast]
	for low <= fast {           // 由于区间定义左闭右闭，所以low=fast有意义
		// 找左边等于val的元素
		for low <= fast && nums[low] != val {
			low++
		}
		// 找右边不等于val的元素
		for fast >= low && nums[fast] == val {
			fast--
		}
		// 将右边不是val的元素覆盖左边是val的元素
		if low < fast { // 注意这里等于号没有意义
			nums[low] = nums[fast]
			low++ // 替换后要更新索引
			fast--
		}
	}
	return low // low一定指向了最终数组末尾的下一个元素（这个要细品：low定义时指向val元素，当结束，它一定指向时不等于val的元素。 ）
}