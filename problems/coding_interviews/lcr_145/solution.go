package lcr145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CheckSymmetricTree 判断二叉树是否对称。
func CheckSymmetricTree(root *TreeNode) bool {
	var recur func(left, right *TreeNode) bool
	recur = func(left, right *TreeNode) bool {
		if left == right {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return recur(left.Left, right.Right) && recur(left.Right, right.Left)
	}
	return root == nil || recur(root.Left, root.Right)
}
