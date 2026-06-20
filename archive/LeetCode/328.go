/*
328. 奇偶链表
2023-3-8
linK: https://leetcode.cn/problems/odd-even-linked-list/
question:
	给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
	第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
	请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
	你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
answer:
	注意时间和空间复杂度，所以需要在原链表上进行修改。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head.Next
	p1, p2 := head, odd // p1, p2分别指向奇数链和偶数链

	for p2 != nil && p2.Next != nil {
		p1.Next = p2.Next
		p2.Next = p2.Next.Next
		p1 = p1.Next
		p2 = p2.Next

	}
	p1.Next = odd
	return head
}