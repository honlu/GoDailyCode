package hot100

/*
40-二叉树的直径
https://leetcode.cn/problems/diameter-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	// 判断每个节点的左右子树高度和, 这一题感觉不像简单题呀
	var res int
	var dfs func(node *TreeNode) int // 深度优先算法
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		lHight := dfs(node.Left) + 1 // 此公式很重要，+1，以及node=nil返回-1，也是很巧妙的公式
		rHight := dfs(node.Right) + 1
		res = max(res, lHight+rHight)
		return max(lHight, rHight)
	}
	dfs(root)
	return res
}
