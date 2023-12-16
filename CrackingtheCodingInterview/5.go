/*
面试题 02.05. 链表求和
2023-4-2 by lu
link: https://leetcode.cn/problems/sum-lists-lcci/
question:
	给定两个用链表表示的整数，每个节点包含一个数位。
	这些数位是反向存放的，也就是个位排在链表首部。
	编写函数对这两个整数求和，并用链表形式返回结果。
answer:
	双指针、虚拟头节点
*/
// 第一遍代码，待改进
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 注意是反向存放的
	head := &ListNode{}
	p := head
	p1, p2 := l1, l2
	flag := 0
	for p1 != nil && p2 != nil {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = (p1.Val + p2.Val + flag) % 10
		flag = (p1.Val + p2.Val + flag) / 10
		p1 = p1.Next
		p2 = p2.Next
	}
	for p1 != nil {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = (p1.Val + flag) % 10
		flag = (p1.Val + flag) / 10
		p1 = p1.Next

	}
	for p2 != nil {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = (p2.Val + flag) % 10
		flag = (p2.Val + flag) / 10
		p2 = p2.Next
	}
	if flag != 0 {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = flag
	}
	return head.Next
}