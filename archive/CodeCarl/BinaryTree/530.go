package binarytree

import "math"

/*
22
530、二叉搜索树的最小绝对差
day：2022-6-10
update: 2023-1-27
link:https://leetcode.cn/problems/minimum-absolute-difference-in-bst/
question:
	给你一个二叉搜索树的根节点 root ，
	返回 树中任意两不同节点值之间的最小差值 。
	差值是一个正数，其数值等于两值之差的绝对值。
idea:
	注意是二叉搜索树，二叉搜索树中序遍历可是有序的。
	最直观的想法，就是把二叉搜索树转换成有序数组，
	然后遍历一遍数组，就统计出来最小差值了。

	同时要学会在递归遍历的过程中如何记录前后两个指针，
	(保存上一个节点，再加上当前节点，就两个节点|指针了。)
	这也是一个小技巧，学会了还是很受用的。
*/

// 递归实现——先遍历整个树，转换成有序数组；然后遍历一遍数组
func getMinimumDifference(root *TreeNode) int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	min := math.MaxInt // 一个很大的值
	for i, j := 1, len(res); i < j; i++ {
		temp := res[i] - res[i-1]
		if temp < min {
			min = temp
		}
	}
	return min
}

// 代码改进：中序遍历的同时计算最小值，技巧就是：要保存上一个节点
func getMinimumDifference(root *TreeNode) int {
	var pre *TreeNode // 保存上一个节点
	var traversal func(node *TreeNode)
	min := math.MaxInt // 一个很大的值
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if pre != nil && node.Val-pre.Val < min {
			min = node.Val - pre.Val
		}
		pre = node
		traversal(node.Right)
	}
	traversal(root)

	return min
}

// 递归能做的，栈也能做！迭代实现。
func getMinimumDifference(root *TreeNode) int {
	// 栈声明和初始化
	stack := []*TreeNode{}
	// 递归遍历节点初始化
	cur := root
	var pre *TreeNode
	min := math.MaxInt
	// 中序迭代遍历
	for cur != nil || len(stack) > 0 {
		if cur != nil { // 指针来访问节点，要一直访问到最底层
			stack = append(stack, cur) // 将访问的节点进栈
			cur = cur.Left             // 左
		} else {
			cur = stack[len(stack)-1] // 栈顶元素，即最底层的节点
			stack = stack[:len(stack)-1]

			if pre != nil && cur.Val-pre.Val < min { // 中
				min = cur.Val - pre.Val
			}

			pre = cur
			cur = cur.Right // 右
		}
	}
	return min
}
