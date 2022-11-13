/*
21. 合并两个有序链表
2022-11-13
link: https://leetcode.cn/problems/merge-two-sorted-lists/
question:
	将两个升序链表合并为一个新的 升序 链表并返回。
	新链表是通过拼接给定的两个链表的所有节点组成的。
answer:
	两种方式：迭代或递归。
	这里采用迭代实现！
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head // 这里
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1 // 这里
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return head.Next // 这里
}