package hot100

/*
29-删除链表的倒数第N个节点
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 双指针
	result := &ListNode{Next: head} // 创建一个虚拟头节点
	p, q := result, result
	i := 0
	for q != nil && i < n {
		q = q.Next
		i++
	}
	for q != nil {
		if q.Next == nil {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
		q = q.Next
	}
	return result.Next
}
