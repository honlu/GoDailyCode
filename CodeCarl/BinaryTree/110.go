package binarytree

/*
12、平衡二叉树
day:2022-6-6
upodate: 2023-1-18
link: https://leetcode.cn/problems/balanced-binary-tree/
question:
	给定一个二叉树，判断它是否是高度平衡的二叉树。
	本题中，一棵高度平衡二叉树定义为：
	一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

idea:
	一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

	二叉树节点的深度：指从根节点到该节点的最长简单路径边的条数。
	二叉树节点的高度：指从该节点到叶子节点的最长简单路径边的条数。

	因为求深度可以从上到下去查 所以需要前序遍历（中左右），
	而高度只能从下到上去查，所以只能后序遍历（左右中）

	本题迭代法其实有点复杂，大家可以有一个思路，也不一定说非要写出来。
	但是递归方式是一定要掌握的！
*/

// 递归实现是否是平衡二叉树，利用后序遍历思想（左右中）
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

// 递归实现树的最大高度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 后面：第二版本代码
/* 注意两个条件：
   1. 每个节点
   2. 左右两个子树的高度差绝对值不超过1
*/
func isBalanced(root *TreeNode) bool {
	// root为空
	if root == nil {
		return true
	}
	// root的两个子节点是平衡
	if isBalanced(root.Left) == false || isBalanced(root.Right) == false {
		return false
	}
	// 判断子节点是否平衡，即左右子树的高度差绝对值不超过1
	// 首先求左右子节点的高度
	left := depth(root.Left)
	right := depth(root.Right)

	return abs(left-right) <= 1
}

func abs(nums int) int {
	if nums >= 0 {
		return nums
	} else {
		return -nums
	}
}

// 迭代实现树的深度，即层次遍历
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 队列初始化
	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res++ // 层数更新
		queue = queue[size:]
	}
	return res
}
