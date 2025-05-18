package hot100

/*
30-两两交换链表中的节点
https://leetcode.cn/problems/swap-nodes-in-pairs/description/?envType=study-plan-v2&envId=top-100-liked
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 需要一个虚拟头结点
	result := &ListNode{Next: head}
	p := result
	flag := 0
	for p.Next != nil && p.Next.Next != nil {
		one := p.Next
		two := p.Next.Next
		p.Next = two
		one.Next = two.Next
		two.Next = one
		if flag == 0 {
			result.Next = p.Next
		}
		flag++
		p = p.Next.Next

	}
	return result.Next
}
