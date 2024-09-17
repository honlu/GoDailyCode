/*
题目：训练计划2
- https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 给定一个头节点为 head 的链表用于记录一系列核心肌群训练项目编号，请查找并返回倒数第 cnt 个训练项目编号。
即
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
- 例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。

题解：
- 双指针解决
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainingPlan(head *ListNode, cnt int) *ListNode {
	// 双指针
	temp, cur := head, head
	for cur != nil && cnt != 0 {
		cur = cur.Next
		cnt--
	}
	for cur != nil {
		temp = temp.Next
		cur = cur.Next
	}
	return temp
}