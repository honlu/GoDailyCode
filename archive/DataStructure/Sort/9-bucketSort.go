/*
9. 桶排序
day: 2022-10-9
描述：
	将值为i的元素放入i号桶，最后一次把桶里的元素倒出来。

算法步骤：
	1、初始状态下，整个序列R[1,2,……,n]处于无序状态，且大小在[min,max]范围内
	2、设置桶的数量为bucketNum，则数据可以划分为[0,bucketNum]、[bucketNum,2*bucketNum-1]、……、[n*(bucketNum-1)/bucketNum,n]，数组中数据分别分配到相应桶中
	3、再对每个非空桶中的元素进行排序[可以选择插入排序算法等]
	4、所有的非空桶依次合并即得到排序好的数据

稳定性：
	稳定
复杂度：
	最好：O(n+k)
	最坏：O(n2)
	平均：O(n+k)
*/
package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 5, 3, 6, 7, 2, 9}
	arr = bucketSort(arr, 9) // 注意这里9是一个桶中存放的数量
	fmt.Println(arr)
}

func bucketSort(arr []int, bucketNum int) []int {
	min, max := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
	}
	// 桶的初始化
	bucketCount := (max-min)/bucketNum + 1
	bucket := make([][]int, bucketCount)

	// 利用映射函数将数据分配到各个桶中
	index := 0
	for i := 0; i < len(arr); i++ {
		index = (arr[i] - min) / bucketNum // 数据分配到桶索引index = (value-min)/bucketNum
		bucket[index] = append(bucket[index], arr[i])
	}
	// 桶内排序，并将排好序的桶数据一次合并到已经拍好序的数组中
	temp := 0
	for i := 0; i < bucketCount; i++ {
		if len(bucket[i]) > 0 {
			bucket[i] = insertionSort(bucket[i]) // 这里选择插入排序算法

			copy(arr[temp:], bucket[i]) // 合并

			temp += len(bucket[i])
		}
	}
	return arr
}

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
