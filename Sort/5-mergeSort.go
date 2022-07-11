package main

import "fmt"

/*
归并排序：典型的分而治之-分治法思想的算法应用。实现有两种：
	1、自上而下的递归（所有递归的方法都可以迭代实现，所以有第二种）
	2、自下而上的迭代。

算法步骤：
	1、申请空间，大小为两个已经排序序列之和，该空间用来存放合并后的序列
	2、设定两个指针，最初分别指向两个已排序序列的起始位置
	3、比较两个指针指向的元素，选择小的元素放入到合并空间，并移动指向小元素的指针到下一个位置
	4、重复步骤3直到某一指针达到末尾
	5、将另一序列剩下的元素直接复制到合并序列的末尾
稳定

复杂度：
最好、最差、平均都是 O(nlogn)
*/

// 递归实现
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		// 分治：两两拆分，一直拆到基本元素才向上递归
		return arr
	}
	mid := length / 2
	left := mergeSort(arr[:mid])  // 左侧数据递归拆分
	right := mergeSort(arr[mid:]) // 右侧数据递归拆分
	result := merge(left, right)  // 排序，合并
	return result
}

func merge(left []int, right []int) []int {
	var result []int
	i, j := 0, 0                  // 初始两个指针
	l, r := len(left), len(right) // 两个已排序序列的长度
	for i < l && j < r {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// 将某个序列剩余的元素，添加到合并序列中
	result = append(result, left[i:]...) // ...使多个参数导入
	result = append(result, right[j:]...)
	return result
}

func main() {
	var arr []int
	arr = []int{1, 2, 5, 8, 4, 9}
	arr = mergeSort(arr)
	fmt.Println(arr)
}
