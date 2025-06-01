package hot100

/*
32-随机链表的复制
https://leetcode.cn/problems/copy-list-with-random-pointer/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil { // 细节，边界
		return head
	}
	// 哈希
	nodeMap := make(map[*Node]*Node)
	p := head
	for p != nil {
		nodeMap[p] = &Node{Val: p.Val, Next: p.Next, Random: p.Random} // 深拷贝
		p = p.Next
	}
	result := nodeMap[head]
	result.Next = nodeMap[head.Next]
	result.Random = nodeMap[head.Random]
	p = result.Next
	for p != nil {
		p.Next = nodeMap[p.Next]
		p.Random = nodeMap[p.Random]
		p = p.Next
	}
	return result
}
