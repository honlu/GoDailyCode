package main

import "fmt"

/*
希尔排序：
描述：每一轮按照实现决定的间隔进行插入排序，
	间隔会依次缩小，最后一次一定是1。
	是基于插入排序的以下两个性质而改进的：
		插入排序在对几乎已经排好序的数据操作时，效率高，既可以达到下线性排序的效率；
		但插入排序一般来说是低效的，因为插入排序每次只能将数据移动一位。


算法步骤：
	1、选择一个增量序列t1,t2,t3...tk,其中ti>tj, tk=1。最简单的增量序列为len/2,len/4,...1
	2、按照增量序列的个数为k,对序列进行k躺排序
	3、每趟排序，根据对应的增量ti,将待排序列恩格成若干长度为m的子序列，分别对各子表进行直接插入排序。
	（仅增量为1时，整个序列作为整体来处理）
不稳定

时间复杂度：
	最好：T(n) = O(nlog n)
	最坏：T(n) = O(nlog2 n)
	平均：T(n) =O(nlog2n)

*/

func shellSort(arr []int) []int {
	length := len(arr)
	// 初始增量gap=length/2,按照增量序列为k个，对序列进行k趟排序
	for gap := length / 2; gap > 0; gap /= 2 {
		// 对每个序列进行插入排序
		for i := gap; i < length; i++ {
			// for j := i - gap; j >= 0; j -= gap { // 便于理解，判断放for里面
			// 	if arr[j] > arr[j+gap] {
			// 		arr[j], arr[j+gap] = arr[j+gap], arr[j]
			// 	}
			// }
			for j := i - gap; j >= 0 && arr[j] > arr[j+gap]; j -= gap { // 优化，早日结束
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
	}
	return arr
}

func main() {
	var arr []int
	arr = []int{1, 2, 5, 8, 4, 9}
	arr = shellSort(arr)
	fmt.Println(arr)
}
