package binarytree

/*
14
404、左叶子之和
day:2022-6-6
update: 2023-1-22
link:https://leetcode.cn/problems/sum-of-left-leaves/
question:
	给定二叉树的根节点 root ，返回所有左叶子之和。
idea:
	首先要注意是判断左叶子，不是二叉树左侧节点，所以不要上来想着层序遍历。
	因为题目中其实没有说清楚左叶子究竟是什么节点，那么我来给出左叶子的明确定义：
		如果左节点不为空，且左节点没有左右孩子，那么这个节点的左节点就是左叶子.

	有时树中不存在左结点：需要通过节点的父节点来判断，
	如果该节点的左节点不为空，该节点的左节点的左节点为空，
	该节点的左节点的右节点为空，则找到了一个左叶子
*/

// 递归实现（深度优先搜索）
func someOfLeftLeaves(root *TreeNode) int {
	res := 0
	var findLeft func(node *TreeNode)
	// 深度优先搜索：前序遍历
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

// 迭代版本：广度优先搜索
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 队列和结果初始化
	res := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 这里面是关键
		size := len(queue)
		for i := 0; i < size; i++ {
			temp := queue[i]
			// 关键：判断左叶子节点(不需要考虑它的相邻右节点情况)
			if temp.Left != nil && temp.Left.Left == nil && temp.Left.Right == nil {
				res += temp.Left.Val
			}

			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		queue = queue[size:]
	}
	return res
}
