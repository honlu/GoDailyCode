package main

// 二叉搜索数的最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 注意是二叉搜索树
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			break
		}
	}
	return root
}
