/*
116. 填充每个节点的下一个右侧节点指针
2022-12-29
link: https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
question:
	给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。
	填充它的每个 next 指针，让这个指针指向其下一个右侧节点。
	如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
	初始状态下，所有 next 指针都被设置为 NULL。

answer：
	将三叉树中的相邻节点抽象成一个节点，这样三叉树就可以想象成二叉树，然后去遍历这个二叉树。
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

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	// 遍历:将三叉树中的相邻节点抽象成一个节点，这样三叉树就可以想象成二叉树，然后去遍历这个二叉树。
	traverse(root.Left, root.Right)

	return root
}

// 这里只需要在一存在的相邻节点连接起来，不需要在意最后找不到下一个右侧节点
func traverse(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	// 将传入的两个节点穿起来
	left.Next = right

	// 继续连接相邻的两个节点
	// 连接相同父节点的两个子节点
	traverse(left.Left, left.Right)
	traverse(right.Left, right.Right)
	// 连接跨越父节点的两个子节点
	traverse(left.Right, right.Left)
}