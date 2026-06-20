package hot100

/*
27-合并两个有序链表
https://leetcode.cn/problems/merge-two-sorted-lists/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 双指针
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var p, q *ListNode
	if list1.Val <= list2.Val {
		p, q = list1, list1
		list1 = list1.Next
	} else {
		p, q = list2, list2
		list2 = list2.Next
	}
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			q.Next = list1
			q = q.Next
			list1 = list1.Next
		} else {
			q.Next = list2
			q = q.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		q.Next = list1
	}
	if list2 != nil {
		q.Next = list2
	}
	return p
}
