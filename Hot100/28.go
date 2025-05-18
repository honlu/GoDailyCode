package hot100

/*
28-两数相加
https://leetcode.cn/problems/add-two-numbers/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, flag int
	result := &ListNode{}
	p := result
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + flag
		flag = sum / 10
		temp := &ListNode{
			Val: sum % 10,
		}
		p.Next = temp
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 != nil {
		l1 = l2
	}
	for l1 != nil {
		sum = l1.Val + flag
		flag = sum / 10
		temp := &ListNode{
			Val: sum % 10,
		}
		p.Next = temp
		p = p.Next
		l1 = l1.Next
	}
	if flag > 0 {
		temp := &ListNode{
			Val: flag,
		}
		p.Next = temp
	}
	return result.Next
}
