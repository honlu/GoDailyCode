package lcr160

import "container/heap"

// 数据流中的中位数
// 注意：有序整数列表
// 这里需要补充一下知识点：优先队列和堆
/*
优先队列是可以根据队列内元素优先级来执行出队操作的数据结构。
可以用堆来实现 O(logN) 时间复杂度入队和出队操作的优先队列，通常优先队列也确实这么实现。

什么是堆？
堆是一个完全二叉树，堆中某个节点的值总是不大于或不小于其父节点的值。
*/
// golang当中堆都要根据类型实现5个API接口
// Less决定是大优先还是小优先
type minHeap []int

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	x := (*m)[(*m).Len()-1]
	*m = (*m)[:(*m).Len()-1]
	return x
}

type maxHeap []int

func (m *maxHeap) Len() int {
	return len(*m)
}

func (m *maxHeap) Less(i, j int) bool { // 大根堆
	return (*m)[i] > (*m)[j]
}
func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}
func (m *maxHeap) Push(item interface{}) {
	(*m) = append((*m), item.(int))
}
func (m *maxHeap) Pop() interface{} {
	item := (*m)[(*m).Len()-1]
	(*m) = (*m)[:(*m).Len()-1]
	return item
}

type MedianFinder struct {
	minH minHeap
	maxH maxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		[]int{},
		[]int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minH.Len() == this.maxH.Len() {
		// 如果两个堆的大小相等，则小顶堆增加元素「增加的不一定是num」
		if this.minH.Len() == 0 || this.minH[0] < num {
			heap.Push(&this.minH, num)
		} else {
			heap.Push(&this.maxH, num)
			top := heap.Pop(&this.maxH).(int)
			heap.Push(&this.minH, top)
		}
	} else {
		if num > this.minH[0] {
			heap.Push(&this.minH, num)
			bottle := heap.Pop(&this.minH).(int)
			heap.Push(&this.maxH, bottle)
		} else {
			heap.Push(&this.maxH, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minH.Len() == this.maxH.Len() {
		return float64(this.minH[0])/2.0 + float64(this.maxH[0])/2.0
	} else {
		return float64(this.minH[0])
	}
}
