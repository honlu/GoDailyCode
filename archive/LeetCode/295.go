/*
295. 数据流的中位数
2023-3-14
link: https://leetcode.cn/problems/find-median-from-data-stream/
question:
	中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，
	中位数是两个中间值的平均值。

	例如 arr = [2,3,4] 的中位数是 3 。
	例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
	实现 MedianFinder 类:
	MedianFinder() 初始化 MedianFinder 对象。
	void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
	double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10-5 以内的答案将被接受。
answer:
	初看挺简单也直接写，就是一个切片的插入、中间值及算法，但提交部分实例超时了！
	再看题解：要使用堆，这种高效抽象的数据结构。
	利用go原有包中heap实现
*/
/*
堆实现：这样时间和空间复杂度才不会O(n)
*/
package main

import (
	"container/heap"
	"containter/heap"
)

// 要使用heap首先要自己定义两个结构体，来实现是大顶堆还是小顶堆
// 小顶堆
type MinHeap struct{
    nums []int
}

func NewMinHeap() *MinHeap{
    return &MinHeap{}
}

func (this *MinHeap) Len() int{
    return len(this.nums)
}

func (this *MinHeap) Swap(i, j int){
    this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

func (this *MinHeap) Less(i, j int) bool{
    return this.nums[i] < this.nums[j]
}

func (this *MinHeap) Push(x interface{}){
	this.nums = append(this.nums, x.(int))
}

func (this *MinHeap) Pop() interface{
    x := this.nums[this.Len()-1]
    this.nums = this.nums[:this.Len()-1]
    return x
}

func (this *MinHeap) Top() int{
    x := this.nums[0]  // 注意，无论是大顶堆还是小顶堆，堆顶都是第一个元素
    return x
}

// 大顶堆
type MaxHeap struct{
    nums []int
}

func NewMaxHeap() *MaxHeap{
    return &MaxHeap{}
}

func (this *MaxHeap) Len() int{
    return len(this.nums)
}

func (this *MaxHeap) Swap(i, j int){
    this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

func (this *MaxHeap) Less(i, j int) bool{
    return this.nums[i] > this.nums[j]   // 大小顶堆仅仅这里不一样。
}

func (this *MaxHeap) Push(x interface{}){
	this.nums = append(this.nums, x.(int))
}

func (this *MaxHeap) Pop() interface{ // 推出最大元素
    x := this.nums[this.Len()-1]
    this.nums = this.nums[:this.Len()-1]
    return x
}

func (this *MaxHeap) Top() int{
    x := this.nums[0]  // 注意，无论是大顶堆还是小顶堆，堆顶都是第一个元素
    return x
}

/* 
将已添加的元素分两部分，前半部分用大顶堆存储，后半部分用小顶堆存储。
大顶堆最多比小顶堆多存储一个元素。
*/
type MedianFinder struct {
	maxHe *MaxHeap
    minHe *MinHeap
}

func Constructor() MedianFinder {
    res := MedianFinder{
		maxHe: NewMaxHeap(),
		minHe: NewMinHeap(),
	}
	// 关键部分，没有太理解
	heap.Init(res.maxHe)
	heap.Init(res.minHe)

	return res
}

func (this *MedianFinder) AddNum(num int)  {
	heap.Push(this.maxHe, num)  // 先加入到大顶堆，这是num可能不是大顶堆里最大的，所以顺序会发生改变

	heap.Push(this.minHe, heap.Pop(this.maxHe)) // 重新调整两个堆的平衡，此时大顶堆Pop出最大元素，加入到小顶堆并且重新排序
	// heap.Pop(this.maxHe) 和this.maxHe.Pop()有什么区别呢？
	for this.minHe.Len() > this.maxHe.Len(){
		heap.Push(this.maxHe, heap.Pop(this.minHe))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHe.Len() > this.minHe.Len(){
		return float64(this.maxHe.Top())
	}else{
		return float64(this.maxHe.Top()+this.minHe.Top())*0.5
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/*
以下是没有通过的代码，仅供提醒！简单的数据结构实现！
*/
type MedianFinder struct {
	nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		nums: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	index := sort.SearchInts(this.nums, num)
	if index < len(this.nums) {
		this.nums = append(this.nums[:index], append([]int{num}, this.nums[index:]...)...)
	} else {
		this.nums = append(this.nums, num)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.nums)
	if n%2 == 0 {
		return (float64(this.nums[n/2-1]) + float64(this.nums[n/2])) / 2
	} else {
		return float64(this.nums[n/2])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
