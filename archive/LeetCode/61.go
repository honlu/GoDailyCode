/*
61. 旋转链表
2023-2-15 by lu
link:https://leetcode.cn/problems/rotate-list/
question:
	给你一个链表的头节点 head ，旋转链表，
	将链表每个节点向右移动 k 个位置。
answer:
	双指针，模拟
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	// 避免k大于链表长度
	size := 0
	p := head
	for p != nil {
		size++
		p = p.Next
	}
	k = k % size
	if k == 0 {
		return head
	}

	p = head
	q := head
	for q.Next != nil {
		for k > 0 {
			q = q.Next
			k--
		}
		if q.Next != nil { // 注意这里！
			p = p.Next
			q = q.Next
		}
	}
	temp := p.Next
	p.Next = nil
	q.Next = head

	return temp
}