package stack

import "math"

/*
70-最小栈
https://leetcode.cn/problems/min-stack/description/?envType=study-plan-v2&envId=top-100-liked
*/
type MinStack struct {
	queue []int
	min   int
}

func Constructor() MinStack {
	return MinStack{min: math.MaxInt}
}

func (this *MinStack) Push(val int) {
	if val < this.min {
		this.min = val
	}
	this.queue = append(this.queue, val)
}

func (this *MinStack) Pop() {
	val := this.Top()
	this.queue = this.queue[:len(this.queue)-1]
	// 如果栈顶是最小元素，重新在栈内选择最小的元素
	if val == this.min && len(this.queue) > 0 {
		this.min = math.MaxInt
		for _, item := range this.queue {
			if item < this.min {
				this.min = item
			}
		}
	} else if len(this.queue) == 0 { // 全出栈，min更新
		this.min = math.MaxInt
	}
}

func (this *MinStack) Top() int {
	return this.queue[len(this.queue)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
