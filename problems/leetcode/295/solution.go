package p295

import "container/heap"

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type MedianFinder struct {
	left  maxHeap
	right minHeap
}

func Constructor() MedianFinder {
	finder := MedianFinder{}
	heap.Init(&finder.left)
	heap.Init(&finder.right)
	return finder
}

func (f *MedianFinder) AddNum(num int) {
	heap.Push(&f.left, num)
	heap.Push(&f.right, heap.Pop(&f.left))
	if f.right.Len() > f.left.Len() {
		heap.Push(&f.left, heap.Pop(&f.right))
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if f.left.Len() > f.right.Len() {
		return float64(f.left[0])
	}
	return float64(f.left[0]+f.right[0]) / 2
}
