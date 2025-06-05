package hot100

/*
46-二叉树展开为链表
https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	// 先序遍历
	var queue []*TreeNode
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		queue = append(queue, node)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	for i := 0; i < len(queue)-1; i++ {
		queue[i].Left = nil
		queue[i].Right = queue[i+1]
	}
}
