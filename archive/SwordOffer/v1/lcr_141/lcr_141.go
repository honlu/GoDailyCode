package lcr141

/*
题目：训练计划3
- https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/?envType=study-plan-v2&envId=coding-interviews
- 给定一个头节点为 head 的单链表用于记录一系列核心肌群训练编号，请将该系列训练编号 倒序 记录于链表并返回。
示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]

示例 3：
输入：head = []
输出：[]

题解：
- 递归
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainningPlan(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 递归终止条件
		return head
	}
	newHead := trainningPlan(head.Next) // 递归到最后一个节点
	head.Next.Next = head               // 反转,细品
	head.Next = nil                     // 断开
	return newHead
}
