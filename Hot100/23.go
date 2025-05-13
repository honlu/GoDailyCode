package hot100

/*
23-反转链表
https://leetcode.cn/problems/reverse-linked-list/description/?envType=study-plan-v2&envId=top-100-liked

*/
// 迭代法解决，有空再写递归式的解法
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	p := head
	var result *ListNode
	for p != nil {
		next := p.Next
		p.Next = result
		result = p
		p = next
	}
	return result
}
