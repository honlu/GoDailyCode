/*
117. 填充每个节点的下一个右侧节点指针2
2022--12-29
link：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
question:
	和116相比，不再是完美二叉树！
	填充它的每个 next 指针，让这个指针指向其下一个右侧节点。
	如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
	初始状态下，所有 next 指针都被设置为 NULL.
answer:
	将两个存在的相邻节点想象成一个节点！
	层次遍历，将每一层的相邻节点连接起来！
*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// 特点：不再是完美二叉树

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		nums := len(queue)
		for i := 0; i < nums; i++ {
			if i == nums-1 { // 两个存在相邻节点连接
				queue[i].Next = nil
			} else {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[nums:]
	}
	return root
}