package hot100

/*
38-反转二叉树
https://leetcode.cn/problems/invert-binary-tree/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	// 递归实现，递归还是要练
	if root != nil {
		root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	}
	return root
}
