package hot100

/*
31-K个一组反转链表
https://leetcode.cn/problems/reverse-nodes-in-k-group/?envType=study-plan-v2&envId=top-100-liked

第一遍没做出来，一时没想出来递归写法，后面再写
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
解题思路：还是回溯，但是有点绕
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	var res *ListNode
	node := head
	// 检查是否有 k 个节点
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	// 反转前 k 个节点（不含 node）
	res = reverse(head, node)
	// 反转后，head 是末尾，接上递归的下一段
	head.Next = reverseKGroup(node, k)
	return res
}

// 实现两个节点之间的链表反转.关键
// 使用两个指针反转 [first, last) 区间
func reverse(first *ListNode, last *ListNode) *ListNode {
	p, q := first, first.Next
	for q != last {
		temp := q.Next
		q.Next = p
		p = q
		q = temp
	}
	return p // 注意返回p，不是q
}
