/*
剑指offer34. 二叉树中和为某一值的路径
2023-4-14 by lu
link: https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
question:
	给你二叉树的根节点 root 和一个整数目标和 targetSum ，
	找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
	叶子节点 是指没有子节点的节点。

answer:
	深度优先搜索，遍历从根节点到叶子节点路径。
	注意：Golang里面，切片是一个细节坑！
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int { // 深度遍历,先序遍历
	res := [][]int{}
	path := []int{}
	var traversal func(node *TreeNode, num int, path []int) // 参数还是要path吗？注意，必须要！！
	traversal = func(node *TreeNode, num int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		num -= node.Val
		// if node.Left == nil && node.Right == nil{
		//     if num == 0{
		//         res = append(res, append([]int{},path...))  // 注意1：这里!!!非常重要
		//     }
		//     return // 注意这里可能导致返回太早，没有回溯。不过path作为参数就可以避免了！
		// }
		if node.Left == nil && node.Right == nil && num == 0 {
			res = append(res, append([]int{}, path...)) // 注意1：这里!!!非常重要

			return
		}
		traversal(node.Left, num, path)
		traversal(node.Right, num, path)
		path = path[:len(path)-1] // 回溯
	}
	traversal(root, target, path)
	return res
}