/*
25. K个一组翻转链表
2022-11-10
link：https://leetcode.cn/problems/reverse-nodes-in-k-group/
question：
	给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
	k 是一个正整数，它的值小于或等于链表的长度。
	如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
	你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
answer:
	双指针，滑动窗口。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归解法
func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	// 对k个一组的链表进行反转
	newHead := reverse(head, cur)     // 当反转后，则head变成此组的最后一个元素
	head.Next = reverseKGroup(cur, k) // 递归解法，容易理解
	return newHead
}

// 返回反转后的头节点
func reverse(start, end *ListNode) *ListNode {
	var pre *ListNode
	cur := start
	for cur != end {
		nx := cur.Next
		cur.Next = pre
		pre = cur
		cur = nx
	}
	return pre
}