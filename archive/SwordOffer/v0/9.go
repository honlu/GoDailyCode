/*
9. 用两个栈实现队列
2022-12-21
link: https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
question:
	用两个栈实现一个队列。队列的声明如下，
	请实现它的两个函数 appendTail 和 deleteHead ，
	分别完成在队列尾部插入整数和在队列头部删除整数的功能。
	(若队列中没有元素，deleteHead 操作返回 -1 )

answer:
	首先要有队列的结构体，以及创建函数，还有队列的插入和删除方法
*/
type CQueue struct {
	stack1 []int // 用于元素添加的栈
	stack2 []int // 用于元素删除的临时栈
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(stack2) > 0 { // 这里关键，当stack2中有元素，一定要删除完后，才能再从stack1中移过来
		temp := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		return temp
	}
	if len(this.stack1) == 0 {
		return -1
	}
	for len(this.stack1) > 0 {
		this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
	temp := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return temp
}