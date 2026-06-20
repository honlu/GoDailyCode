package lcr136

import "testing"

// 单元测试
func TestDeleteNode(t *testing.T) {
	// 构建链表
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1 := &ListNode{Val: 1}
	node9 := &ListNode{Val: 9}
	node4.Next = node5
	node5.Next = node1
	node1.Next = node9

	// 删除节点
	newHead := deleteNode(node4, 5)
	if newHead.Val != 4 || newHead.Next.Val != 1 || newHead.Next.Next.Val != 9 {
		t.Fatal("测试失败")
	}
}

// 单元测试
func TestDeleteNode2(t *testing.T) {
	// 构建链表,链表中有重复值
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1 := &ListNode{Val: 1}
	node5_2 := &ListNode{Val: 5}
	node9 := &ListNode{Val: 9}
	node4.Next = node5
	node5.Next = node1
	node1.Next = node5_2
	node5_2.Next = node9

	// 删除节点
	newHead := deleteNode2(node4, 5)
	if newHead.Val != 4 || newHead.Next.Val != 1 || newHead.Next.Next.Val != 9 {
		t.Fatal("测试失败")
	}
}
