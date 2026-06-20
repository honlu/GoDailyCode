package heap

import "container/heap"

/*
76-295.数据流的中位数
https://leetcode.cn/problems/find-median-from-data-stream/?envType=study-plan-v2&envId=top-100-liked

*/

// 以下简单添加，排序。超时。
// type MedianFinder struct {
// 	queue []int
// }

// func Constructor() MedianFinder {
// 	return MedianFinder{}
// }

// func (this *MedianFinder) AddNum(num int) {
// 	this.queue = append(this.queue, num)
// 	sort.Ints(this.queue)

// }

// func (this *MedianFinder) FindMedian() float64 {
// 	if len(this.queue)%2 == 0 {
// 		return float64(this.queue[(len(this.queue)/2)-1]+this.queue[len(this.queue)/2]) * 1.0 / 2.0
// 	} else {
// 		return float64(this.queue[len(this.queue)/2])
// 	}
// }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/*
上述排序算法是全局重排，复杂度较高。而堆的话，是局部调整，相对复杂较低。大顶堆和小顶堆只需要保证堆顶元素，不需要切片所有有序。

 思路：达到当前算法的抽象思维和实践能力的边界。
 对顶堆，左半边是大顶堆，右半边是小顶堆。中位数只和堆顶有关。
这里假设左顶堆等于或比左顶堆多一个.

另外加强理解：什么是堆？以及堆的元素加入和删除细节流程。
*/
// 小顶堆实现
type MiniHeap []int

func (c MiniHeap) Len() int            { return len(c) }
func (c MiniHeap) Swap(i, j int)       { c[i], c[j] = c[j], c[i] }
func (c MiniHeap) Less(i, j int) bool  { return c[i] < c[j] }       // 小的优先
func (c *MiniHeap) Push(x interface{}) { *c = append(*c, x.(int)) } // 添加元素
func (c *MiniHeap) Pop() interface{} {
	old := *c // 为什么还要添加一个中间变量呢？
	count := len(old)
	x := old[count-1]
	*c = old[:count-1]
	return x
}

// 大顶堆实现
type MaxHeap []int

func (c MaxHeap) Len() int            { return len(c) }
func (c MaxHeap) Swap(i, j int)       { c[i], c[j] = c[j], c[i] }
func (c MaxHeap) Less(i, j int) bool  { return c[i] > c[j] }       // 大的优先
func (c *MaxHeap) Push(x interface{}) { *c = append(*c, x.(int)) } // 添加元素
func (c *MaxHeap) Pop() interface{} {
	old := *c
	count := len(old)
	x := old[count-1]
	*c = old[:count-1]
	return x
}

type MedianFinder struct {
	left  *MaxHeap  // 左边大顶堆，默认如果两边元素已经相等时，再添加元素到大顶堆
	right *MiniHeap // 右边小顶堆
}

/*
思路：
我们维护两个堆：
- 最大堆（左半部分）：存较小的一半，堆顶是左边的最大值。
- 最小堆（右半部分）：存较大的一半，堆顶是右边的最小值。

插入规则：
- 新数 ≤ 左堆堆顶 → 插左堆
- 否则 → 插右堆

保持平衡：两个堆的大小差 ≤ 1。

中位数：
如果长度相等 → (左堆顶 + 右堆顶) / 2
如果不相等 → 元素更多的那个堆的堆顶
*/
func Constructor() MedianFinder {
	l, r := &MaxHeap{}, &MiniHeap{}
	heap.Init(l)
	heap.Init(r)
	return MedianFinder{l, r}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == this.right.Len() {
		heap.Push(this.right, num) // 注意和直接使用this.right.Push(num)有着非常的不同
		v := heap.Pop(this.right)
		heap.Push(this.left, v)
	} else {
		heap.Push(this.left, num)
		v := heap.Pop(this.left)
		heap.Push(this.right, v)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		return float64((*this.left)[0]+(*this.right)[0]) / 2.0
	}
	return float64((*this.left)[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
