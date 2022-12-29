package binarytree

/*
27.修剪二叉搜索树
day:2022-6-15
link:https://leetcode.cn/problems/trim-a-binary-search-tree/
idea:
在剪枝的时候，可以分为三步：
将root移动到[L, R] 范围内，注意是左闭右闭区间
剪枝左子树
剪枝右子树
*/

// 递归实现-剪枝二叉搜索树
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
