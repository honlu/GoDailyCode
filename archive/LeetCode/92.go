/*
92. 反转链表2
2022-11-18
link: https://leetcode.cn/problems/reverse-linked-list-ii/
question:
	给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
	请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表
answer:
	注意：指定片段内的反转！
	模拟操作链表
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if right == left { // 当两者相等时，不需要反转
		return head
	}
	fast := head
	low := &ListNode{}
	low.Next = head
	flag := 1
	for fast != nil && flag != left {
		fast = fast.Next
		flag++
		low = low.Next
	}
	// 开始反转
	p := fast
	temp := fast.Next
	for fast != nil && temp != nil && flag != right {
		fast = temp // 1 注意这里不能写成fast = fast.Next,否则会造成反向了！
		temp = temp.Next
		fast.Next = p
		p = fast
		flag++
	}
	// 开始拼接
	p = low.Next
	low.Next = fast
	p.Next = temp
	// 这里有些边界注意,即left=1导致的特殊情况
	if p == head {
		return fast
	} else {
		return head
	}
}