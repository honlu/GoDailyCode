package hot100

/*
22-相交链表
https://leetcode.cn/problems/intersection-of-two-linked-lists/?envType=study-plan-v2&envId=top-100-liked
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
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
			continue // 注意这里到nil时要跳过
		}
		if b == nil {
			b = headA
			continue
		}
		a = a.Next
		b = b.Next
	}
	return a
}
