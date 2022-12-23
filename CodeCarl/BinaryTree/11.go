package binarytree

/*
12、平衡二叉树
day:2022-6-6
idea:一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

二叉树节点的深度：指从根节点到该节点的最长简单路径边的条数。
二叉树节点的高度：指从该节点到叶子节点的最长简单路径边的条数。

因为求深度可以从上到下去查 所以需要前序遍历（中左右），
而高度只能从下到上去查，所以只能后序遍历（左右中）

本题迭代法其实有点复杂，大家可以有一个思路，也不一定说非要写出来。
但是递归方式是一定要掌握的！
*/

// 迭代实现是否是平衡二叉树，利用后序遍历思想（左右中）
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) { // 左右
		return false
	}
	leftHeight := maxDepth(root.Left) + 1 // maxDepth递归实现树的深度，在8.go中
	rightHeight := maxDepth(root.Right) + 1

	return abs(leftHeight-rightHeight) <= 1 // 代码优化
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
