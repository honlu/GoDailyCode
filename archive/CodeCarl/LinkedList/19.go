/*
19. 删除链表中的倒数第N个节点
2023-1-6
link: https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
question:
	给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
answer:
	双指针
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 设置虚拟头节点，方便删除头节点
	p := &ListNode{}
	p.Next = head
	low, fast := p, p.Next
	// 让fast先跑n步
	for n > 0 {
		fast = fast.Next
		n--
	}
	// 两者移动到最后
	for fast != nil {
		low = low.Next
		fast = fast.Next
	}
	low.Next = low.Next.Next // 当fast为nil,low.Next指向倒数第n个节点
	return p.Next
}