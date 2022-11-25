/*
2.两数相加
2022-11-25
link: https://leetcode.cn/problems/add-two-numbers/
question:
	给你两个非空的链表，表示两个非负的整数。
	它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
	请你将两个数相加，并以相同形式返回一个表示和的链表。
answer:
	注意：链表形式的两数相加。注意Next的使用！
	模拟操作时，注意细节！
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := l1   // 在l1指向的链表上进行更改，用来保存结果链表
	head := l3 // head最终返回的
	flag := 0  // 进位标志
	for l1 != nil && l2 != nil {
		temp := l1.Val + l2.Val + flag
		if temp >= 10 {
			temp %= 10
			flag = 1
		} else {
			flag = 0
		}
		l1.Val = temp
		l2 = l2.Next
		l3 = l1
		l1 = l1.Next

	}
	if l2 != nil {
		l3.Next = l2
	}
	for l3.Next != nil {
		l3 = l3.Next // 注意这里和上面的位置不同
		temp := l3.Val + flag
		if temp >= 10 {
			temp %= 10
			flag = 1
		} else {
			flag = 0
		}
		l3.Val = temp
	}
	if flag == 1 {
		temp := &ListNode{1, nil}
		l3.Next = temp
	}
	return head
}