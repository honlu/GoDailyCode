/*
86. 分割链表
2023-4-16 by lu
link: https://leetcode.cn/problems/partition-list/
question:
	给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，
	使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

	你应当 保留 两个分区中每个节点的初始相对位置。
answer:
	注意链表的操作：双指针，可以设置两个链表，也可以使用一个虚拟节点进行修改。
*/
// 个人模拟，虚拟头节点实现，有些复杂
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil { // 这个不能少
		return head
	}
	p1 := &ListNode{}
	res := p1
	p1.Next = head
	var left, right *ListNode
	for p1.Next != nil {
		if p1.Next.Val < x {
			p1 = p1.Next
		} else {
			left = p1
			right = p1.Next
			break
		}
	}
	if right != nil { // 有可能上面else里面执行不到，right没有赋值
		p1 = right
	}
	for p1.Next != nil {
		if p1.Next.Val < x {
			temp := p1.Next
			p1.Next = p1.Next.Next
			left.Next = temp
			left = left.Next
			temp.Next = right
		} else {
			p1 = p1.Next
		}
	}
	return res.Next
}

// 参考题解优化：双链表优化！
func partition(head *ListNode, x int) *ListNode {
	min, max := &ListNode{}, &ListNode{} // 小链表，大于或等于链表
	left, right := min, max
	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}
	right.Next = nil // 这个不能少！
	left.Next = max.Next
	return min.Next
}