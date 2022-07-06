package binarytree

/*
25. 二叉搜索树中的插入操作
day:2022-6-15
link:https://leetcode.cn/problems/insert-into-a-binary-search-tree/
idea:
可以不考虑题目中提示所说的改变树的结构的插入方式，
只要按照二叉搜索树的规则去遍历，遇到空节点就插入节点就可以了。
*/

// 递归遍历二叉搜索树，遇到空节点就插入节点
// func insertIntoBST(root *TreeNode, val int) *TreeNode {
// 	if root == nil{
// 		root = &TreeNode{Val: val}
// 		return root
// 	}
// 	if root.Val > val{
// 		root.Left = insertIntoBST(root.Left, val)
// 	}else{
// 		root.Right = insertIntoBST(root.Right, val)
// 	}
// 	return root
// }

// 迭代实现-或者成为·模拟·，参看官方答案
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	p := root // 模拟节点遍历
	for p != nil {
		if p.Val > val {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				break
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				break
			}
			p = p.Right
		}
	}
	return root
}
