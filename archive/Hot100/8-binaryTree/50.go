package hot100

import "math"

/*
50-124.二叉树中的最大路径和
https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 本题难点在于：发现问题解决点，转到求每个节点的路径和，以及如何涉及实现
func maxPathSum(root *TreeNode) int {
	// 遍历每个节点的最大贡献值
	res := math.MinInt
	// 后序遍历，遍历每个节点返回当前节点的最大路径和
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0) // 如果子节点的贡献值小于0，则路径和不需要子节点
		// 节点的最大路径和：取决于该节点的值和左右子节点的最大贡献值
		res = max(res, node.Val+left+right)
		// 返回节点的最大贡献值
		return node.Val + max(left, right)
	}
	dfs(root)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
