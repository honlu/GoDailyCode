/*
堆排序
day:2022-10-9
描述：(最大堆，有序区)
	从堆顶把根卸出来放在有序区之前，再恢复堆。
	堆是一个近似完全二叉树的结构，并满足子节点的键值或索引总是小于或大于它的父节点。
	堆排序可以说是一种利用堆的概念来排序的选择排序。
	每个节点的值都大于或等于其子节点的值，为最大堆；反之为最小堆。

	在堆的算法里，最大堆用一个数组来表示。在数组表示中，第 i 个元素的子节点位于 2*i+1 和 2*i+2。

算法步骤：
	1、创建一个堆R[0,..., n-1],此时是无序堆，而堆顶是最大元素。
	2、再将堆顶R[0]和无序区的最后一个记录R[n-1]交换，
		由此得到新的无序区R[0…n-2]和有序区R[n-1]，
		且满足R[0…n-2].keys ≤ R[n-1].key
	3、由于交换后新的根R[1]可能违反堆性质，故应将当前无序区R[1..n-1]调整为堆。
		然后再次将R[1..n-1]中关键字最大的记录R[1]和该区间的最后一个记录R[n-1]交换，
		由此得到新的无序区R[1..n-2]和有序区R[n-1..n]，
		且仍满足关系R[1..n-2].keys≤R[n-1..n].keys，
		同样要将R[1..n-2]调整为堆。
	4、重复步骤2，直到无序区只有一个元素为止。

稳定性：
	不稳定
复杂度：
	最好：O(nlogn)
	最坏：O(nlogn)
	平均：O(nlogn)

*/
package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 5, 3, 6, 7, 2, 9}
	arr = heapSort(arr) // 注意这里使用的是len(arr)-1，不是len(arr)
	fmt.Println(arr)
}

// 堆排序
func heapSort(arr []int) []int {
	// 1、构建堆(这里用大顶堆构建升序)
	// 2、调整堆，把堆顶元素和第i-1个元素交换，
	// 这样0....i-2就又成为一个堆，继续对这个堆进行构建，调整
	buildMaxHeap(arr, len(arr)) // 先构建n个元素的大顶堆
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i] // 调整堆顶元素，把堆顶元素和最后一个元素交换
		buildMaxHeap(arr, i)
	}
	return arr
}

// 建立最大堆
// // 构建堆，一般从最后一个非叶子节点开始构建，即从下往上调整，从下往上能让最大（小）值元素转移到堆顶
func buildMaxHeap(arr []int, cap int) {
	// 非叶子节点的i范围从0...(n/2-1)个
	for i := (cap / 2) - 1; i >= 0; i-- {
		// 调整左子树
		left := 2*i + 1
		if left < cap && arr[i] < arr[left] {
			arr[i], arr[left] = arr[left], arr[i]
		}
		// 调整右子树
		right := 2*i + 2
		if right < cap && arr[i] < arr[right] {
			arr[i], arr[right] = arr[right], arr[i]
		}
	}
}
