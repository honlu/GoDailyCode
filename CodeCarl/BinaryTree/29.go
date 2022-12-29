package binarytree

/*
29.把二叉搜索树转换为累加树
day:2022-6-15
link:https://leetcode.cn/problems/convert-bst-to-greater-tree/
idea:
从树中可以看出累加的顺序是右中左，所以我们需要
反中序遍历这个二叉树，然后顺序累加就可以了。
*/

// 递归实现反中序遍历，进而产生类累加树
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var traversal func(*TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil { // 终止条件，遇到空节点，就返回
			return
		}
		// 非空节点时，就计算。按照右中左即反中序遍历
		traversal(node.Right) // 右节点
		// 中节点处理
		sum += node.Val
		node.Val = sum
		traversal(node.Left) // 左结点
	}
	traversal(root)
	return root
}
