/*
707. 设计链表
link: https://leetcode.cn/problems/design-linked-list/
question:
	设计链表的实现。您可以选择使用单链表或双链表。
	单链表中的节点应该具有两个属性：val 和 next。
	val 是当前节点的值，next 是指向下一个节点的指针/引用。
	如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。
	假设链表中的所有节点都是 0-index 的。

	在链表类中实现这些功能：
		get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
		addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。
			插入后，新节点将成为链表的第一个节点。
		addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
		addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。
			如果 index 等于链表的长度，则该节点将附加到链表的末尾。
			如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
		deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

answer:
	设计型题目，面试很容易问类似的系统设计题目。
	链表有两种：单链表和双链表。

["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail"
,"get","addAtHead","addAtIndex","addAtHead"]
[[],[7],[2],[1],[3,0],[2],[6],[4]
,[4],[4],[5,0],[6]]
*/
package main

// 单链表设计
// 单链表节点定义
type Node struct { // 注意LeetCode不能定义成ListNode，否则会出现序列化错误
	Val  int
	Next *Node
}

// 个人单链表类定义
type MyLinkedList struct {
	head *Node
	Size int // 这个Size非常方便索引的判断. Size是从0开始计数的，注意！
}

func Constructor() MyLinkedList {
	return MyLinkedList{&Node{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size { // 索引失效
		return -1
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Next.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val) // 注意this.head是一个虚拟头节点
}

func (this *MyLinkedList) AddAtTail(val int) {
	// 因为自己定义有Size，就非常方便在末尾添加
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	cur := this.head // 注意这里head，是一个虚拟头节点
	for i := 0; i < index; i++ {
		cur = cur.Next
	} // 注意，for循环结束，cur指向
	node := &Node{Val: val}
	node.Next = cur.Next
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size { // 索引无效
		return
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next // 由于this.head是虚拟头节点，所以删除头节点即index=0时，可以直接删和其他节点删除一样
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
