/*
54. 二叉搜索树的第K大节点
2022-12-22
link: https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
question:
	给定一棵二叉搜索树，请找出其中第 k 大的节点的值。
answer:
	注意：二叉搜索树的特点：左孩子节点比父节点小，右孩子节点比父节点大。
	中序遍历是有序的！（从小到大）
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	var res []int
	midTraversal(root, &res)
	return res[len(res)-k]
}

// 中序遍历
func midTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	midTraversal(root.Left, res)
	*res = append(*res, root.Val)
	midTraversal(root.Right, res)
}