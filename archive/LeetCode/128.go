/*
128. 复制带随机指针的链表
2022-12-15
link: https://leetcode.cn/problems/copy-list-with-random-pointer/
question:
	给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，
	该指针可以指向链表中的任何节点或空节点。
	构造这个链表的 深拷贝。
answer:
	难点：只知道原链表的头节点head.如果新建节点，如何让随机节点指向成功呢？
	这里使用map哈希来解决指向问题,map的键原链表节点，值为新链表节点
*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	hash := make(map[*Node]*Node)
	// 组建hash
	p := head
	for p != nil {
		hash[p] = &Node{Val: p.Val} // 哈希
		p = p.Next
	}
	// 建立新链表
	p = head
	for p != nil {
		hash[p].Next = hash[p.Next]
		hash[p].Random = hash[p.Random]
		p = p.Next
	}
	return hash[head]
}