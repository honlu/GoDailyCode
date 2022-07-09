package main

import "fmt"

/*
选择排序：
（有序区，无序区）
在无序区里找一个最小的元素跟在有序区的后面。对数组：比较的多，换的少
	每次在未排序序列中选择最小元素放到最前面。

算法步骤：
	1、首先在未排序序列中找到最小（大）元素，存放在排序序列的起始位置
	2、再从剩余未排序元素继续寻找最小（大）元素，放到已排序序列的末尾
	3、重复2，直到所有元素均排序完毕

对数组不稳定：举例，a=b>c，且在待排序序列中a,b,c依次前后顺序；由于c和a会进行选择交换，则a.b相对顺序就破坏了。

时间复杂度：
	最佳：O(n^2)
	最差：O(n^2)
	平均：O(n^2)

什么时候最好或最坏？
	没有，因为算法本身就是什么数据都是O(n^2)的复杂度，所以用它的时候，数据规模越小越好。
*/

func selectionSort(arr []int) []int {
	length := len(arr)
	// 第一层循环：代表已排序的长度
	for i := 0; i < length-1; i++ {
		min := i // 保存最小元素的索引
		// 第二层循环：未排序的序列开始
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	var arr []int
	arr = []int{1, 5, 3, 6}
	arr = selectionSort(arr)
	fmt.Println(arr)
}
