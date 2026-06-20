package lcr123

/*
题目：图书整理1
- https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 书店店员有一张链表形式的书单，每个节点代表一本书，节点中的值表示书的编号。
  为更方便整理书架，店员需要将书单倒过来排列，就可以从最后一本书开始整理，逐一将书放回到书架上。请倒序返回这个书单链表。

示例 1：
输入：head = [3,6,4,1]
输出：[1,4,6,3]

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

func reverseBookList(head *ListNode) []int {
	var res []int
	var reverse func(head *ListNode)
	reverse = func(head *ListNode) {
		if head == nil { // 递归终止条件
			return
		}
		reverse(head.Next)          // 递归
		res = append(res, head.Val) // 递归后处理
	}
	reverse(head)
	return res
}
