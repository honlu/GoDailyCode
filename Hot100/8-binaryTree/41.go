pakcage hot100

/* 
41-二叉树的层次遍历
https://leetcode.cn/problems/binary-tree-level-order-traversal/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var queue []*TreeNode
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		var temp []int
		tLen := len(queue)
		for i := 0; i < tLen; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, temp)
		queue = queue[tLen:]
	}
	return res
}