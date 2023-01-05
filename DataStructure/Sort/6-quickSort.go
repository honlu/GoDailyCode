/*
快速排序
day:2022-10-9
描述：(小数，基准元素，大数)
	在区间内随机挑选一个元素作为基准，将小于基准的元素放在基准之前，
	大于基准的元素放在基准之后，再分别对小数区和大数区进行排序。

算法步骤：
	1、从数列中跳出一个元素，称为“基准”pivot
	2、重新排列数列，比基准值小的元素放在其前，大的放在其后（相同数可以放任一边）。
		在这个分区退出之后，该基准就处于数列的中间位置。此步成为“partition”分区操作。
	3、递归地把小于基准值元素的子序列和大于基准值的子序列进行排序。

稳定性：
	不稳定
复杂度：
	最好：O(nlogn)
	最坏：O(n2)
	平均：O(nlogn)

注意事项：关于边界范围很重要！
*/
package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 5, 3, 6, 7, 2, 9}
	arr = quickSort(arr, 0, len(arr)-1) // 注意这里使用的是len(arr)-1，不是len(arr)
	fmt.Println(arr)
}

// 快排
func quickSort(arr []int, left, right int) []int { // 注意区间定义[left,right]
	if left < right {
		index := parition(arr, left, right) // 返回索引
		// 递归
		quickSort(arr, left, index-1)
		quickSort(arr, index+1, right)
	}
	return arr
}

// 分区
func parition(arr []int, left, right int) int {
	privot := left                    // 随机选择一个基准，就按照最左的
	index := privot + 1               // 分区索引
	for i := index; i <= right; i++ { // 划分分区，将[left+1,index]区域都小于基准，后面都大于基准
		if arr[i] < arr[privot] { // 注意这里和基准比较
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}
	arr[privot], arr[index-1] = arr[index-1], arr[privot]

	return index - 1
}
