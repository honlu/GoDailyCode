package binarytree

/*
10、完全二叉树的节点个数
day:2022-6-5
idea:分外普通二叉树解法和利用完全二叉树特性的递归解法
*/

// 普通二叉树，求有多少节点，直接无脑遍历相加算长度就行
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	return res
}

// 利用完全二叉树特性的递归解法
func countNodesSpecial(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := 0, 0
	// 直接通过最左侧的树节点和最右侧的树节点判断是否是满二叉树
	left, right := root.Left, root.Right
	for left != nil {
		left = left.Left
		leftHeight += 1
	}
	for right != nil {
		right = right.Right
		rightHeight += 1
	}
	if leftHeight == rightHeight {
		return (2 << leftHeight) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1 // 注意这个递归，如果某个节点左右子树不是满二叉树，就会+1，并且一直下去
}
