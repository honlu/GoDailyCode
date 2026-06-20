/*
剑指offer52. 两个链表的第一个公共交点
2023-4-4 by lu
link: https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
question:
	输入两个链表，找出它们的第一个公共节点。
answer:
	两个链表循环遍历，当同时走相同长度时，就可以判断是否是相交点。注意两个链表可能不存在相交点。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil { // 注意这里不能使用p1.Next == nil来判断，否则如果两个链表不存在相交点会死循环
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}