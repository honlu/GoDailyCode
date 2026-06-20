package lcr171

/*
题目：训练计划v
- https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 输入两个链表，找出它们的第一个公共节点。
- 如下面的两个链表：
- headA: a1 -> a2 -> c1 -> c2 -> c3
- headB: b1 -> b2 -> b3 -> c1 -> c2 -> c3
- 在节点 c1 开始相交。

题解：
- 双指针解决
- 两个链表长度分别为 m 和 n，m > n
- 1. 先让长链表的指针先走 m - n 步
- 2. 然后两个链表的指针一起走，直到相遇

时间复杂度分析：
- 时间复杂度：O(m + n)
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { // 如果有一个链表为空，则直接返回 nil
		return nil
	}
	p, q := headA, headB // 双指针
	for p != q {         // 如果两个链表没有相交，则 p 和 q 最终会同时指向 nil
		p = p.Next
		q = q.Next
		if p == q { // 如果两个链表相交，则直接返回 p
			return p
		}
		if p == nil { // 如果 p 为空，则让 p 指向 headB
			p = headB
		}
		if q == nil { // 如果 q 为空，则让 q 指向 headA
			q = headA
		}
	}
	return p
}
