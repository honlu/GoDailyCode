package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var firstOrder func(a, b *TreeNode) bool
	firstOrder = func(a, b *TreeNode) bool {
		// 前序遍历
		if b == nil {
			return true
		}
		if a == nil || a.Val != b.Val {
			return false
		}
		return firstOrder(a.Left, b.Left) && firstOrder(a.Right, b.Right)
	}
	return (A != nil && B != nil) && (firstOrder(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}
