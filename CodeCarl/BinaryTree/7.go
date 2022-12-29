package binarytree

/*
7、对称二叉树
day:2022-6-5
idea:首先想清楚，判断对称二叉树要比较的是哪两个节点，要比较的可不是左右节点！
对于二叉树是否对称，要比较的是根节点的左子树与右子树是不是相互翻转的，理解这一点就知道了其实我们要比较的是两个树（这两个树是根节点的左右子树），所以在递归遍历的过程中，也是要同时遍历两棵树。
*/

// 递归-实现:
// 递归函数
func compare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return compare(left.Left, right.Right) && compare(right.Left, left.Right)
}

func isSymmetric(root *TreeNode) bool { // 判断是否是对称二叉树
	return compare(root.Left, root.Right)
}

// 迭代实现
func isSymmetricIter(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]  // 要比较的左结点
		right := queue[1] // 要比较的右节点
		queue = queue[2:] // 剩余的结点
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left) // 重点：最开始两个要对应起来
	}
	return true
}
