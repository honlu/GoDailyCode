package main

// LCR 145 - 判断对称二叉树
func checkSymmetricTree(root *TreeNode) bool {
	// 递归
	var recur func(left, right *TreeNode) bool
	recur = func(left, right *TreeNode) bool {
		if left == right { // 两个都为nil
			return true
		} else if left == nil || right == nil || left.Val != right.Val { // 其他
			return false
		} else { // 两个都不为nil且值相等
			return recur(left.Left, right.Right) && recur(left.Right, right.Left)
		}

	}
	return root == nil || recur(root.Left, root.Right)
}
