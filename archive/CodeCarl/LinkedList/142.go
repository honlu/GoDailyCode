/*
142. 环形链表2
2023-1-7
link: https://leetcode.cn/problems/linked-list-cycle-ii/
question:
	给定一个链表的头节点 head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
	如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
	为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
	如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
	不允许修改 链表。
answer:
	两种方式：
	1. 暴力hash存储节点，进而遍历可以判断是否成环。
	2. 进阶：使用O(1)空间复杂度解决
		定义 fast 和 slow 指针，从头结点出发，fast指针每次移动两个节点，slow指针每次移动一个节点，
		如果 fast 和 slow指针在途中相遇 ，说明这个链表有环。并且一定是在环中相遇。

*/
// 暴力法
func detectCycle(head *ListNode) *ListNode {
	hash := make(map[*ListNode]bool) // 哈希存储
	p := head
	for hash[p] != true && p != nil { // 节点已经存储并且不为nil时，就接触遍历
		hash[p] = true
		p = p.Next
	}
	return p
}

