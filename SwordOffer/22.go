/*
22. 链表中倒数第k个节点
2022-11-29
link:
	https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
question:
	输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，
	本题从1开始计数，即链表的尾节点是倒数第1个节点。
answer:
	标准的快慢指针（双指针）题目！
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k <= 0 { // 处理k小于等于0情况
		return nil
	}
	low, fast := head, head
	for k != 0 {
		k--
		fast = fast.Next
		if fast == nil {
			return head // 处理k大于链表长度情况
		}
	}
	for fast != nil {
		fast = fast.Next
		low = low.Next
	}
	return low
}