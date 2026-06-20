package lcr143

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsSubStructure 判断 B 是否为 A 的子结构。
func IsSubStructure(a *TreeNode, b *TreeNode) bool {
	var firstOrder func(a, b *TreeNode) bool
	firstOrder = func(a, b *TreeNode) bool {
		if b == nil {
			return true
		}
		if a == nil || a.Val != b.Val {
			return false
		}
		return firstOrder(a.Left, b.Left) && firstOrder(a.Right, b.Right)
	}
	return (a != nil && b != nil) && (firstOrder(a, b) || IsSubStructure(a.Left, b) || IsSubStructure(a.Right, b))
}
