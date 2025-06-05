package hot100

/*
47-从前序与中序遍历序列构造二叉树
https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 中序和先序的特点，从中序中找先序的根节点，然后划分左右子树
	// 注意没有重复元素，是关键
	res := new(TreeNode)
	if len(preorder) == 0 {
		return nil
	}
	res.Val = preorder[0]
	index := find(inorder, preorder[0])
	res.Left = buildTree(preorder[1:1+index], inorder[:index])
	res.Right = buildTree(preorder[1+index:], inorder[index+1:])
	return res
}

func find(order []int, target int) int {
	for i, num := range order {
		if num == target {
			return i
		}
	}
	return -1
}
