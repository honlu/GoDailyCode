/*
225. 用队列实现栈
2023-1-12
link: https://leetcode.cn/problems/implement-stack-using-queues/
question:
	请你仅使用两个队列实现一个后入先出（LIFO）的栈，
	并支持普通栈的全部四种操作（push、top、pop 和 empty）。

	实现 MyStack 类：
		void push(int x) 将元素 x 压入栈顶。
		int pop() 移除并返回栈顶元素。
		int top() 返回栈顶元素。
		boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。

	注意：
		你只能使用队列的基本操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
		你所使用的语言也许不支持队列。
		你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。

answer:
	Go中没有队列的数据结构。实现切片来实现队列。
	用两个队列queue和backups实现队列的功能，
	backups其实完全就是一个备份的作用，把queue最后面的元素以外的元素都备份到backups，
	然后弹出最后面的元素，再把其他元素从backups导回queue。
*/
type MyStack struct {
	queue   []int // 队列，用来栈的入和出
	backups []int // 备份队列
}

func Constructor() MyStack {
	return MyStack{
		queue:   make([]int, 0),
		backups: make([]int, 0),
	}
}

// 入栈
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x) // 队尾插入
}

// 出栈(关键)
func (this *MyStack) Pop() int {
	// 将queue只留最后一个元素，其他元素全部到backup中
	for len(this.queue) > 1 {
		this.backups = append(this.backups, this.queue[0])
		this.queue = this.queue[1:]
	}
	// 留下的最后一个元素就是要返回的值
	temp := this.queue[0]
	this.queue = this.queue[1:] // 出队，queue变为空
	// 交换queue和backups
	this.queue, this.backups = this.backups, this.queue
	return temp
}

// 获取栈顶元素
func (this *MyStack) Top() int {
	val := this.Pop()
	this.queue = append(this.queue, val) // 注意：top不出栈，所以要复原
	return val
}

// 判断是否栈是否为空：并让queue只保留top元素，backups备份前面的元素
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */