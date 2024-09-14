package lcr136

/*
题目：删除链表的节点
- https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
- 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。

示例 1:
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]

示例 2:
输入: head = [4,5,1,9], val = 1
输出: [4,5,9]

解题思路：
- 由于给定的是单向链表，无法直接删除节点，只能通过修改节点的值来模拟删除节点的操作。
- 从头节点开始遍历链表，找到要删除的节点，将其值替换为下一个节点的值，然后删除下一个节点。
- 特殊情况：如果要删除的节点是头节点，则直接删除头节点即可。
- 双指针解决
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	pre, cur := head, head.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}
		pre, cur = cur, cur.Next
	}
	return head
}

// 题目变形：删除链表的节点。注意：删除的节点可能有重复值
func deleteNode2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}
