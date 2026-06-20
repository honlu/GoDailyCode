package main

// 二叉树的最近公共祖先
func lowestCommonAncestor194(root, p, q *TreeNode) *TreeNode {
	// 递归如何实现
	if root == nil || root == p || root == q {
		return root
	}
	// 后续遍历
	left := lowestCommonAncestor194(root.Left, p, q)
	right := lowestCommonAncestor194(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil { // 两个节点一定存在于给定的二叉树中
		return left
	} else if right != nil {
		return right
	}
	return nil
}
