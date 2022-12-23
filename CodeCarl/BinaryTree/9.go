package binarytree

import "container/list"

/*
9、二叉树的最小深度
day:2022-6-5
idea:最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
求二叉树的最小深度和求二叉树的最大深度的差别主要在于处理左右孩子不为空的逻辑。
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 递归实现
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil { // 当一个左子树为空，右不为空，这时并不是最低点
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil && root.Left != nil { // 当一个右子树为空，左不为空，这时并不是最低点
		return 1 + minDepth(root.Left)
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// 迭代实现-层次遍历，但要注意当遍历某一层出现一个结点的左右孩子都不存在时，就可以退出了
func minDepthIter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		res += 1
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return res
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}
