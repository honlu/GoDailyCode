/*
143. 重排链表
2022-11-20
link: https://leetcode.cn/problems/reorder-list/
question:
	给定一个单链表L的头节点head,将其重新排列，使其后面节点依次插入到前面两个节点的中间。
answer:
	双指针，递归！
	递归三要素：
		1.确定递归函数的参数和返回值
		2.确定终止条件
		3.确定底层递归的逻辑
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	cur := head
	if cur.Next == nil || cur.Next.Next == nil { // 起始判断，是否可以进行重排
		return
	}
	pre := &ListNode{} // 注意节点的声明和初始化!
	pre.Next = head
	lat := head
	// 找到最后一个节点
	for lat.Next != nil {
		pre = pre.Next
		lat = lat.Next
	}
	// 将最后一个节点插入到第一个和第二个节点当中
	temp := cur.Next
	cur.Next = lat
	pre.Next = nil
	lat.Next = temp
	// 准备递归
	cur = cur.Next.Next
	if cur.Next == pre || cur == pre { // 重排后，在判断是否还需要递归执行重排
		return
	}
	reorderList(cur) // 递归！
}
