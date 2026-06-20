package hot100

/*
26-环形链表2
https://leetcode.cn/problems/linked-list-cycle-ii/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	itemMap := make(map[*ListNode]bool)
	p := head
	for p != nil {
		if itemMap[p] == true {
			return p
		}
		itemMap[p] = true
		p = p.Next
	}
	return nil
}
