package hot100

/*
36-二叉树的中序遍历

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// 递归实现,注意res不能当作递归遍历函数的参数
	var res []int
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}
	traversal(root)
	return res
}
