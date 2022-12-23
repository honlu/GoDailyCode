/*
8. 计数排序
day: 2022-10-9
描述：
	统计小于等于该元素值的元素的个数i, 于是该元素就放在目标数组的索引i为（i>=0）
	计数排序要求输入的数据必须是有确定范围的整数。

算法步骤：
	1、找出待排序的数组中最大和最小的元素
	2、统计数组中每个值为i的元素出现的次数，存入数组C的第i项
	3、对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
	4、反向填充目标数组：将每个元素i放在新数组的第C[i]项，每放一个元素就将C[i]减去1
稳定性：
	稳定
复杂度：
	最好：O(n+k)
	最坏：O(n+k)
	平均：O(n+k)
注意：当输入的元素是 n 个 0 到 k 之间的整数时，它的运行时间是 Θ(n + k)。
计数排序不是比较排序，排序的速度快于任何比较排序算法。
*/
package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 5, 3, 6, 7, 2, 9}
	arr = countSort(arr, 9) // 注意这里9是数组中最大值
	fmt.Println(arr)
}

func countSort(arr []int, max int) []int {
	bucket := make([]int, max+1) // 创建长度为max+1的数组
	// 统计每个值为i的次数
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	// 遍历max个值的数，看其在原来的数组中有多少个，然后填充到原先的数组中
	index := 0 // 要填充的数组索引
	for i := 0; i <= max; i++ {
		for bucket[i] > 0 { // i元素出现次数，然后填充到原来数组中
			arr[index] = i
			bucket[i]-- // 填充一次后，次数减少1
			index++     // 填充索引往后移
		}
	}
	return arr
}
