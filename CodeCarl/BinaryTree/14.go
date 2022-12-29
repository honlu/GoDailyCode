package binarytree

/*
14、左叶子之和
day:2022-6-6
link:https://leetcode.cn/problems/sum-of-left-leaves/
idea:首先要注意是判断左叶子，不是二叉树左侧节点，所以不要上来想着层序遍历。
因为题目中其实没有说清楚左叶子究竟是什么节点，那么我来给出左叶子的明确定义：
如果左节点不为空，且左节点没有左右孩子，那么这个节点的左节点就是左叶子.

有时树中不存在左结点：需要通过节点的父节点来判断，
如果该节点的左节点不为空，该节点的左节点的左节点为空，
该节点的左节点的右节点为空，则找到了一个左叶子
*/

// 递归实现
func someOfLeftLeaves(root *TreeNode) int {
	res := 0
	var findLeft func(node *TreeNode)
	findLeft = func(node *TreeNode) {
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil { // 左叶子节点
			res += node.Left.Val
		}
		if node.Left != nil {
			findLeft(node.Left)
		}
		if node.Right != nil {
			findLeft(node.Right)
		}
	}
	findLeft(root) // 调用
	return res
}
