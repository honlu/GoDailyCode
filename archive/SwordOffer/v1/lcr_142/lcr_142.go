package lcr142

/*
题目：训练计划iv
- https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/?envType=study-plan-v2&envId=coding-interviews
- 给定两个升序链表，请将它们合并为一个新的 升序 链表并返回。
- 新链表是通过拼接给定的两个链表的所有节点组成的。
-
- 示例1：
- 输入：1->2->4, 1->3->4
- 输出：1->1->2->3->4->4
-
- 题解：
- - 递归解决
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlanIV(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = trainingPlanIV(l1.Next, l2)
		return l1
	} else {
		l2.Next = trainingPlanIV(l1, l2.Next)
		return l2
	}
}
