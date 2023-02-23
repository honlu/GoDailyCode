package binarytree

/*
26
450. 删除二叉搜索树中的节点
day:2022-6-15
update:2023-1-28
link:https://leetcode.cn/problems/delete-node-in-a-bst/
question:
	给定一个二叉搜索树的根节点 root 和一个值 key，
	除二叉搜索树中的 key 对应的节点，
	并保证二叉搜索树的性质不变。
	返回二叉搜索树（有可能被更新）的根节点的引用。

	一般来说，删除节点可分为两个步骤：
		首先找到需要删除的节点；
		如果找到了，删除它。

idea:
首先找到需要删除的节点； 如果找到了，怎么删除它？

二叉搜索树中删除节点遇到的情况都搞清楚。
	有以下五种情况：
	第一种情况：没找到删除的节点，遍历到空节点直接返回了

	找到删除的节点
	第二种情况：左右孩子都为空（叶子节点），直接删除节点，
		返回NULL为根节点
	第三种情况：删除节点的左孩子为空，右孩子不为空，删除节点，右孩子补位，
		返回右孩子为根节点
	第四种情况：删除节点的右孩子为空，左孩子不为空，删除节点，左孩子补位，
		返回左孩子为根节点
	（第五种最难）
	第五种情况：左右孩子节点都不为空，则将其右子树的最小节点替换删除节点，
	将其指向删除节点的左子树，以及删除最小节点的右子树
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归 if版
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key { // 找到了进行删除
		// 把情况3和4正确处理
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 处理情况5
		// 获得右子树最小的节点
		node := root.Right
		for node.Left != nil {
			node = node.Left
		}
		// 删除右子树最小的节点
		root.Right = deleteNode(root.Right, node.Val)
		// 用右子树最小的节点替换root节点
		node.Left = root.Left
		node.Right = root.Right
		root = node

	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

// 递归实现 switch版
func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil: // 代表没有搜索到值为 key 的节点，返回空
		return nil
	case root.Val > key: // 表示key的节点可能存在与root的左子树中，递归
		root.Left = deleteNode(root.Left, key)
	case root.Val < key: // 表示key的节点可能存在与root的右子树中，递归
		root.Right = deleteNode(root.Right, key)
	case root.Right == nil || root.Left == nil: // root.val = key，root为删除的节点，但其左右子树有可能为空
		if root.Left != nil { // 只有左子树无右，此时就将左子树作为新的子树，返回其左子节点
			return root.Left
		}
		return root.Right // // 只有右子树无左，此时就将右子树作为新的子树，返回其右子节点
	default: // root.Val = key,并且左右子树都存在！ 难点！
		// 首先找到右子树的最小节点即root 的右子节点，再不停地往左子节点寻找，直到找到一个不存在左子节点的节点
		// ，作为新的root节点，并将其在右子树的位置删除
		node := root.Right
		for node.Left != nil {
			node = node.Left
		}
		node.Right = deleteNode(root.Right, node.Val)
		node.Left = root.Left
		return node
	}
	return root // 这个需要吗？需要！
}
