package binarytree

/*
25
701. 二叉搜索树中的插入操作
day:2022-6-15
update:2023-1-25
link:https://leetcode.cn/problems/insert-into-a-binary-search-tree/
question:
	给定二叉搜索树（BST）的根节点 root 和要插入树中的值 value ，
	将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。
	输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。

	注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。
	你可以返回 任意有效的结果 。

idea:
	不要考虑题目中提示所说的改变树的结构的插入方式，
	只要按照二叉搜索树的规则去遍历，遇到空节点就插入节点就可以了。
*/

// 递归遍历二叉搜索树，遇到空节点就插入节点
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 迭代实现-或者成为·模拟·，参看官方答案
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	pre := root
	// 根据二叉搜索树迭代搜索空节点
	for root != nil {
		if root.Val > val {
			if root.Left == nil { // 空节点
				root.Left = &TreeNode{Val: val}
				break
			}
			root = root.Left // 非空节点
		} else { // root.Val > val
			if root.Right == nil {
				root.Right = &TreeNode{Val: val}
				break
			}
			root = root.Right
		}
	}
	return pre
}
