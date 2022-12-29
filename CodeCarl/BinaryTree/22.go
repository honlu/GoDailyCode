package binarytree

import "math"

/*
22、二叉搜索树的最小绝对差
day：2022-6-10
link:https://leetcode.cn/problems/minimum-absolute-difference-in-bst/
idea:注意是二叉搜索树，二叉搜索树中序遍历可是有序的。
最直观的想法，就是把二叉搜索树转换成有序数组，然后遍历一遍数组，就统计出来最小差值了。

同时要学会在递归遍历的过程中如何记录前后两个指针，这也是一个小技巧，学会了还是很受用的。
*/

// 递归实现——先遍历整个树，转换成有序数组；然后遍历一遍数组
// func getMinimumDifference(root *TreeNode) int {
// 	var res []int
// 	var traversal func(node *TreeNode)
// 	traversal = func(node *TreeNode) {
// 		if node == nil{
// 			return
// 		}
// 		traversal(node.Left)
// 		res = append(res, node.Val)
// 		traversal(node.Right)
// 	}
// 	traversal(root)
// 	min := math.MaxInt // 一个很大的值
// 	for i, j := 1, len(res); i < j; i++{
// 		temp := res[i] - res[i-1]
// 		if temp < min{
// 			min = temp
// 		}
// 	}
// 	return min
// }

// 改进：中序遍历的同时计算最小值，技巧就是：要保存上一个节点
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
