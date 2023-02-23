package binarytree

/*
29
538.把二叉搜索树转换为累加树
day:2022-6-15
update:2023-1-30
link:https://leetcode.cn/problems/convert-bst-to-greater-tree/
question:
	给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），
	使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
idea:
	从树中可以看出累加的顺序是右中左，所以我们需要
	反中序遍历这个二叉树，然后顺序累加就可以了。

	注意:提醒一下，二叉搜索树满足下列约束条件：
	节点的左子树仅包含键 小于 节点键的节点。
	节点的右子树仅包含键 大于 节点键的节点。
	 左右子树也必须是二叉搜索树。
*/

// 递归实现反中序遍历，进而产生类累加树
// 重点提前设计一个变量sum
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

// 迭代，用栈实现反中序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	// 前一个节点的数值
	pre := 0
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil { // （首先把所有右节点加入栈中，然后通过出栈遍历当前节点以及其的左节点）
			stack = append(stack, cur)
			cur = cur.Right // 右
		} else {
			cur = stack[len(stack)-1] // 中
			stack = stack[:len(stack)-1]
			cur.Val += pre
			pre = cur.Val
			cur = cur.Left // 左
		}
	}
	return root
}
