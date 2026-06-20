package lcr147

import (
	"fmt"
	"math"
)

/*
题目：最小栈
- https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，以返回栈元素中的最小值。执行push、pop或min操作均摊时间复杂度O(1)。
*/

type MinStack struct {
	stack []int // 栈：先进后出
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
	return
}

func (this *MinStack) Top() int {
	fmt.Println(this.stack)
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	var min int = math.MaxInt32 // 注意初始值要足够大
	for _, item := range this.stack {
		if min > item {
			min = item
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
