/*
912. 排序数组
link: https://leetcode.cn/problems/sort-an-array/
question:
	给你一个整数数组 nums，请你将该数组升序排列
answer:
	自己写一个排序算法，或者调库。
	以下排序算法是单线程的快速排序。(手写的快排会有用例超时!
	啊：为什么呀?)
	关于一定数量的goroutine并行的快速排序，可参考下面链接：
	https://colobu.com/2018/06/26/implement-quick-sort-in-golang

*/
func sortArray(nums []int) []int {
	// sort.Ints(nums) // 调库
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 快速排序
// func quickSort(nums []int, left, right int) {
// 	if left >= right {
// 		return
// 	}

// 	index := parition(nums, left, right) // 返回基准所在的索引
// 	// 递归
// 	quickSort(nums, left, index-1)  // 小于基准的子序列排序
// 	quickSort(nums, index+1, right) // 大于基准的子序列排序

// }

// // 分区
// func parition(arr []int, left, right int) int {
// 	privot := arr[right]            // 随机选择最左的作为基准
// 	index := left - 1               // 设置分区索引
// 	for i := left; i < right; i++ { // 划分分区，将[left+1,index]区域都小于基准，后面的区域都大于基准
// 		if arr[i] < privot {
// 			index++
// 			arr[i], arr[index] = arr[index], arr[i]
// 		}
// 	}
// 	arr[right], arr[index+1] = arr[index+1], arr[right] // 关键边界，不能少！

// 	return index + 1
// }

// 上面快排有问题.改进，可以通过leetcode
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	privot := arr[left]
	j, index := left+1, right
	for j <= index {
		if arr[j] > privot {
			arr[j], arr[index] = arr[index], arr[j]
			index--
		} else {
			i++
		}
	}
	arr[left], arr[index] = arr[index], arr[left]

	quickSort(arr, left, index-1)
	quickSort(arr, index+1, right)
}