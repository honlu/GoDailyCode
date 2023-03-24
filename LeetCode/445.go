/*
445. 两数相加2
2023-3-18 by lu
link: https://leetcode.cn/problems/add-two-numbers-ii/submissions/
question:
	给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。
	它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
	你可以假设除了数字 0 之外，这两个数字都不会以零开头。

answer:
	方法一：链表反转.
	首先链表是从高位到低位组成的，所以如果两个链表相边遍历便相加，要先反转。
	（下面是方法一的迭代版，也可以改成递归版，下次更新。）
	方法二：遍历保存所有链表的节点值，利用栈思想再相加
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 反转
	r1 := reverse(l1)
	r2 := reverse(l2)
	// 准备相加
	temp := &ListNode{}
	head := temp
	var flag int
	for r1 != nil && r2 != nil {
		sum := r1.Val + r2.Val + flag
		flag = sum / 10
		sum = sum % 10
		temp.Next = &ListNode{Val: sum}
		temp = temp.Next
		r1 = r1.Next
		r2 = r2.Next
	}
	for r1 != nil {
		sum := r1.Val + flag
		flag = sum / 10
		sum = sum % 10
		temp.Next = &ListNode{Val: sum}
		temp = temp.Next
		r1 = r1.Next
	}
	for r2 != nil {
		sum := r2.Val + flag
		flag = sum / 10
		sum = sum % 10
		temp.Next = &ListNode{Val: sum}
		temp = temp.Next
		r2 = r2.Next
	}
	if flag != 0 { // 这个不能少呀
		temp.Next = &ListNode{Val: flag}
	}
	// 相加后的链表反转
	r3 := reverse(head.Next)
	return r3
}

// 链表反转
func reverse(root *ListNode) *ListNode {
	var head *ListNode // nil节点
	for root != nil {
		temp := root.Next
		root.Next = head
		head = root
		root = temp
	}
	return head
}