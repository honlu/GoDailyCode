/*
232. 用栈实现队列
2023-1-12
link: https://leetcode.cn/problems/implement-queue-using-stacks/
question:
	请你仅使用两个栈实现先入先出队列。
	队列应当支持一般队列支持的所有操作（push、pop、peek、empty）.

	实现 MyQueue 类：
		void push(int x) 将元素 x 推到队列的末尾
		int pop() 从队列的开头移除并返回元素
		int peek() 返回队列开头的元素
		boolean empty() 如果队列为空，返回 true ；否则，返回 false
	说明：
	你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
	你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。

answer:
	Golang中没有栈stack这个原生数据结构。
	但可以通过切片or数组或者list(双链表)的部分操作来达到stack操作的目的。
*/
// 使用切片模拟栈，两个切片实现队列
type MyQueue struct {
	stackIn  []int // 输入栈
	stackOut []int // 输出栈
}

// 初始化队列
func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0), // 注意这个逗号不能少
	}
}

// 进队
func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

// 出队
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	temp := this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[:len(this.stackOut)-1]
	return temp
}

// 获取头部元素
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	return this.stackOut[len(this.stackOut)-1]
}

// 队列是否为空，并保证输出栈有元素
func (this *MyQueue) Empty() bool {
	if len(this.stackIn) == 0 && len(this.stackOut) == 0 {
		return true
	}
	// 公用代码抽出来，避免Pop和Peek再执行一次
	if len(this.stackOut) == 0 {
		for len(this.stackIn) > 0 {
			this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
		}
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */