/*
剑指27.二叉树的镜像
2023-3-20
link: https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/
question:
	请完成一个函数，输入一个二叉树，该函数输出它的镜像
answer:
	关键在于每个节点的左右子树交换
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	var traversal func(root *TreeNode)
	// 前序遍历
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return root
}