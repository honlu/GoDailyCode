package hot100

/*
37-二叉树的最大深度
https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	// 二叉树的层次遍历
	var res int
	queue := []*TreeNode{}
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		res++
		levels := len(queue)
		for i := 0; i < levels; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[levels:]
	}
	return res
}
