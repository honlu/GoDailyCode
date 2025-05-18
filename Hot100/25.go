package hot100

/*
25-环形链表
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	itemMap := make(map[*ListNode]bool)
	p := head
	for p != nil {
		if itemMap[p] == true {
			return true
		}
		itemMap[p] = true
		p = p.Next
	}
	return false
}
