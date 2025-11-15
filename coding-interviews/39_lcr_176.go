package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	// 递归左右子树的深度
	var depth func(node *TreeNode) int
	var res bool = true
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		if abs(l-r) > 1 {
			res = false
			goto end
		}
		return max(l, r) + 1
	end:
		return 0
	}
	depth(root)
	return res
}
