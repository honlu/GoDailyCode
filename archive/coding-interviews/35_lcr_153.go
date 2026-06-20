package main

// LCR 153. 二叉树中和为目标值的路径
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathTarget(root *TreeNode, target int) [][]int {
	// 递归和回溯
	var res [][]int
	var path []int
	var dfs func(node *TreeNode, target int)
	dfs = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Val == target && node.Left == nil && node.Right == nil {
			// 下面两行，比 append([]int{}, path...) 更高效和省空间
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		if node.Left != nil {
			dfs(node.Left, target-node.Val)
			path = path[:len(path)-1] // 回溯
		}
		if node.Right != nil {
			dfs(node.Right, target-node.Val)
			path = path[:len(path)-1]
		}
	}
	dfs(root, target)
	return res
}
