package hot100

/*
34-23. 合并 K 个升序链表
https://leetcode.cn/problems/merge-k-sorted-lists/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	// // 比较简单的方式，循环调用两个有序链表合并
	// res := lists[0]
	// for i := 1; i < len(lists); i++ {
	// 	if lists[i] == nil {
	// 		continue
	// 	}
	// 	res = mergeList(res, lists[i])
	// }
	// 还可以再优化，先分治合并前半部分，再合并后半部分
	if m == 1 {
		return lists[0]
	}
	left, right := mergeKLists(lists[:m/2]), mergeKLists(lists[m/2:])
	return mergeList(left, right)
}

// 合并两个有序链表
func mergeList(a, b *ListNode) *ListNode {
	p := new(ListNode)
	res := p
	p1, p2 := a, b
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
			p = p.Next
		} else {
			p.Next = p2
			p2 = p2.Next
			p = p.Next
		}
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return res.Next
}
