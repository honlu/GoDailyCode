package lcr194

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LowestCommonAncestor 返回二叉树中 p 和 q 的最近公共祖先。
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 保留原实现思路：后序遍历，左右子树分别命中时当前节点就是答案。
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
