package binarytree

/*
21:验证二叉搜索树
day:2022-6-9
link:https://leetcode.cn/problems/validate-binary-search-tree/
idea:
要知道中序遍历下，输出的二叉搜索树节点的数值是有序序列。
有了这个特性，验证二叉搜索树，就相当于变成了判断一个序列是不是递增的了。
*/

// 递归中序遍历解法，利用二叉搜索树的中序遍历是有序的，所以设置一个上一个节点，就可以依次判断了
// func isValidBST(root *TreeNode) bool {
// 	if root == nil { // 二叉搜索树也可以是空树
// 		return true
// 	}
// 	var pre *TreeNode // 保存上一个指针，初始为nil，这里出现了问题
// 	// 中序遍历：左中右
// 	left := isValidBST(root.Left) // 左
// 	// 中
// 	if pre != nil && pre.Val >= root.Val { // 如果pre是nil，即刚开始，继续执行；若pre已经保存一个节点，但比后一个节点大，则不是二叉搜索树
// 		return false
// 	}
// 	pre = root

// 	right := isValidBST(root.Right)
// 	return left && right
// } 错误代码
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var traversal func(node *TreeNode) bool
	traversal = func(node *TreeNode) bool {
		if node == nil { // 二叉搜索树也可以是空树
			return true
		}

		// 左
		left := traversal(node.Left) // 注意这里全局变量pre也会进入
		// 中
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		pre = node
		// 右
		right := traversal(node.Right)
		return left && right
	}
	return traversal(root)
}
