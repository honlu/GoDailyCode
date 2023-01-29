package binarytree

/*
27
669.修剪二叉搜索树
day:2022-6-15
update:2023-1-29
link:https://leetcode.cn/problems/trim-a-binary-search-tree/
question:
	给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。
	通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L) 。你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。
idea:
	注意二叉搜索树是有序的，所以不符合的删除很方便。
	递归法
	迭代法：在剪枝的时候，可以分为三步：
		将root移动到[L, R] 范围内，注意是左闭右闭区间
		剪枝左子树
		剪枝右子树
	要注意考虑小于low的右子树。
*/

// 递归实现-剪枝二叉搜索树（前序遍历）
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low { // 如果该节点小于最小值，则该节点更换为该节点的右节点值
		root = trimBST(root.Right, low, high)
		return root
	}
	if root.Val > high { // 如果该节点大于最大值，则该节点更换为该节点的左节点值
		root = trimBST(root.Left, low, high)
		return root
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
