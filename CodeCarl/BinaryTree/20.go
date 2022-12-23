package binarytree

/*
20、二叉搜索书中的搜索
link:https://leetcode.cn/problems/search-in-a-binary-search-tree/
idea:
二叉搜索树是一个有序树：
若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
它的左、右子树也分别为二叉搜索树
这就决定了，二叉搜索树，递归遍历和迭代遍历和普通二叉树都不一样。
*/

// 递归实现， 前序遍历寻找节点
func searchBST(root *TreeNode, val int) *TreeNode {
	// 递归结束条件
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val) // return 不能少
	}
	return searchBST(root.Right, val) // return不能少
}
