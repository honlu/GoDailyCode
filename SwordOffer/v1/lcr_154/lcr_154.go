package lcr154

/*
题目：复杂链表的复制
- https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 请实现 copyRandomList 函数，复制一个复杂链表。
在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
还有一个 random 指针指向链表中的任意节点或者 null。

示例 1：
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

示例 2：
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]

示例 3：
输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]

示例 4：
输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。

解题思路：
- 哈希表
- 拼接+拆分
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)

	// 利用哈希map中存的是（原节点->新节点）的映射关系，新的节点都是独立的没有串联
	for cur := head; cur != nil; {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	// 将新节点串起来，组成新链表
	for cur := head; cur != nil; {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]

}
