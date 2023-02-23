/*
160. 相交链表
2023-1-7
link: https://leetcode.cn/problems/intersection-of-two-linked-lists/
question:
	给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
	如果两个链表不存在相交节点，返回 null 。
answer:
	难点: 用两个指针 p1 和 p2 分别在两条链表上前进，并不能同时走到公共节点，也就无法得到相交节点 c1。
	关键：通过某些方式，让 p1 和 p2 能够同时到达相交节点 c1。
	解决：
		用两个指针 p1 和 p2 分别在两条链表上前进，我们可以让 p1 遍历完链表 A 之后开始遍历链表 B，
		让 p2 遍历完链表 B 之后开始遍历链表 A，这样相当于「逻辑上」两条链表接在了一起。
		如果这样进行拼接，就可以让 p1 和 p2 同时进入公共部分，也就是同时到达相交节点 c1.
		O(m+n),O(1)
	也可以暴力，使用哈希解决。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB //  p1 指向 A 链表头结点，p2 指向 B 链表头结点
	for p1 != p2 {
		if p1 == nil { // p1 走一步，如果走到 A 链表末尾，转到 B 链表
			p1 = headB
		} else { // 注意这里一定要在else里面走，如果不加判断默认都走，一些案例可能会死循环。
			p1 = p1.Next
		}
		if p2 == nil { // p2 走一步，如果走到 B 链表末尾，转到 A 链表
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}