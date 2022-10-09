package main

import "fmt"

/*
插入排序：又叫直接插入排序。
描述：(有序区，无序区)
	把无序区的第一个元素插入到有序区的合适的位置。
	对数组：比较的少，换的多
	原理是通过构建有序序列，对于未排序数据，在已排序序列中从后往前扫描，
	找到相应位置并插入。

算法步骤：
	1、将待排序序列第一个元素看作一个有序序列，将第二个元素到最后一个元素当成是未排序序列
	2、从头到尾扫描未排序序列，将扫描到的每一个元素插入到有序序列的适当位置。
	注意：（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
稳定

时间复杂度：
	最好：O(n)
	最坏：O(n^2)
	平均：O(n^2)

*/

func insertionSort(arr []int) []int {
	length := len(arr)
	// 第一层循环代表：将1-length的数组看作待排序的序列
	for i := 1; i < length; i++ {
		// 第二层循环：将i插入到合适的位置
		// for j := i - 1; j >= 0; j-- { // 此简单，好理解，还可以优化
		// 	if arr[j] > arr[j+1] { // 升序
		// 		arr[j], arr[j+1] = arr[j+1], arr[j]
		// 	}
		// }
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- { // 优化，不需要全部遍历已排序，到了相应位置就中断
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return arr
}

func main() {
	var arr []int
	arr = []int{1, 5, 3, 6}
	arr = insertionSort(arr)
	fmt.Println(arr)
}
