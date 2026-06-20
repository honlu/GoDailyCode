/*
24. 两两交换链表中的节点
2023-1-6
link: https://leetcode.cn/problems/swap-nodes-in-pairs/
question:
	给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
	你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
answer:
	只交换节点内的值吗？双指针实现
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre, cur := head, head.Next
	for pre != nil && cur != nil {
		pre.Val, cur.Val = cur.Val, pre.Val
		if cur.Next != nil && cur.Next.Next != nil {
			pre = cur.Next
			cur = cur.Next.Next
		} else {
			break
		}
	}
	return head
}