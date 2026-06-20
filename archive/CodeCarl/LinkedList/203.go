/*
203.移除链表元素
2023-1-5
link: https://leetcode.cn/problems/remove-linked-list-elements/
question:
	给你一个链表的头节点 head 和一个整数 val ，
	请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
answer:
	Go中带自动垃圾回收机制，所以不需要手动管理内存。
	关于链表操作有两种：
	1. 直接使用原来的链表来进行删除操作。
	2. 设置一个虚拟头结点再进行删除操作。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 关于第一种链表操作：直接使用原来的链表
// 注意如果需要移除头节点，它和移除其他节点的操作方式不同，需要单独写一段逻辑代码。
// func removeElements(head *ListNode, val int) *ListNode {
// 	for head != nil && head.Val == val { // 移除符合条件的头节点。
// 		head = head.Next
// 	}
// 	// 移除其他符合条件的节点
// 	p := head
// 	for p != nil && p.Next != nil {
// 		if p.Next.Val == val {
// 			temp := p.Next.Next  // 删除p.Next节点
// 			p.Next = temp
// 		}else{
// 			p = p.Next
// 		}
// 	}
// 	return head
// }

// 关于链表的第二种操作：设置一个虚拟头结点
func removeElements(head *ListNode, val int) *ListNode {
	// 设置一个虚拟头节点
	p := &ListNode{}
	p.Next = head
	cur := p // 再设置一个移动节点
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next // 删除cur.Next
		} else {
			cur = cur.Next
		}
	}
	return p.Next
}