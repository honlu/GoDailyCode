package hot100

/*
39-对称二叉树

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	// 递归实现
	var res bool
	var recursion func(left, right *TreeNode) bool
	recursion = func(left, right *TreeNode) bool {
		if left != nil && right != nil && left.Val == right.Val {
			return recursion(left.Left, right.Right) && recursion(left.Right, right.Left)
		} else if left == nil && right == nil {
			return true
		} else {
			return false
		}
	}
	res = recursion(root.Left, root.Right)
	return res
}
