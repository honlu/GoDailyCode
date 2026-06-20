package p7linkedlist33

import "sort"

/*
33-排序链表
https://leetcode.cn/problems/sort-list/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
 一般sort就可以解决。但是面试官可能不会让你这样做🤔
*/
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 两次for循环
	var temp []int
	p := head
	for p != nil {
		temp = append(temp, p.Val)
		p = p.Next
	}
	sort.Ints(temp)
	res := &ListNode{
		Val: temp[0],
	}
	p = res
	for i := 1; i < len(temp); i++ {
		temp := &ListNode{
			Val: temp[i],
		}
		p.Next = temp
		p = p.Next
	}
	return res
}
