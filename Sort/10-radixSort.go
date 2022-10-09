/*
10. 基础排序
day: 2022-10-9
描述：
	一种多关键字的排序算法，可用桶排序实现。
	将整数按位数切割成不同的数字，然后按每个位数分别比较。

算法步骤：
	1、将所有待比较数值同意为同样的数位长度，数位较短的数前面补零
	2、从最低位开始，一次进行一次排序。
		这样从最低位排序一直到最高位排序完成以后，数列就变成一个有序序列。
稳定性：
	稳定
复杂度：
	最好：O(n*k)
	最坏：O(n*k)
	平均：O(n*k)
注意：
	基数排序是非比较型排序算法。
补充：
	基数排序 vs 计数排序 vs 桶排序
	这三种排序算法都利用了桶的概念，但对桶的使用方法上有明显差异：
		基数排序：根据键值的每位数字来分配桶；
		计数排序：每个桶只存储单一键值；
		桶排序：每个桶存储一定范围的数值
*/

package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 5, 3, 6, 7, 2, 9}
	arr = radixSort(arr)
	fmt.Println(arr)
}

func radixSort(arr []int) []int {
	max := 0 // 最大值
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	// 获取最大值的位数
	count := 0
	for max%10 > 0 {
		max /= 10
		count++
	}
	// 给桶中的位置放数据
	temp := make([]int, len(arr)) // 存放每一位次排序后的结果
	bucket := new([10]int)        // ?
	radix := 1
	var i, j, k int
	for i = 0; i < count; i++ { // 进行count次排序
		for j = 0; j < 10; j++ { // 初始化
			bucket[j] = 0
		}
		for j = 0; j < len(arr); j++ {
			k = (arr[j] / radix) % 10 // 这个很有趣呀！
			bucket[k]++
		}
		for j = 1; j < 10; j++ { //将tmp中的为准依次分配给每个桶
			bucket[j] = bucket[j-1] + bucket[j]
		}
		for j = len(arr) - 1; j >= 0; j-- {
			k = (arr[j] / radix) % 10
			temp[bucket[k]-1] = arr[j]
			bucket[k]--
		}
		for j = 0; j < len(arr); j++ {
			arr[j] = temp[j]
		}
		radix = radix * 10
	}
	return arr
}
