package binarytree

/*
2、二叉树的递归遍历
Day:2022-6-3
*/
// 前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode) // 匿名函数
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val) // 前序添加元素
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val) // 中序添加元素
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// 后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) { // 递归函数
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val) // 后序添加元素
	}
	traversal(root) // 开始执行递归函数
	return res
}
