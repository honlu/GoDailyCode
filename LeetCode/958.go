/*
958. 二叉树的完全性检验
2022-12-23
link: https://leetcode.cn/problems/check-completeness-of-a-binary-tree/
question:
	给定一个二叉树的 root ，确定它是否是一个 完全二叉树 。
	在一个 完全二叉树 中，除了最后一个关卡外，所有关卡都是完全被填满的，
	并且最后一个关卡中的所有节点都是尽可能靠左的。
	它可以包含 1 到 2h 节点之间的最后一级 h 。

answer:
	队列实现层次遍历。[退出条件是关键！]
	[在层序遍历一个二叉树的时候，一个非空节点之前不能有空节点.]
	用一个布尔类型的变量empty来记录在层序遍历的过程中，是否有遇到非空节点。
	如果遇到了，则将empty置为true，在后续遍历一个非空节点时，判断empty是否为true，如果是true则说明不是满二叉树!
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode // 这里可以扩展一个知识点*[]TreeNode和[]*TreeNode区别
	queue = append(queue, root)
	empty := false
	for len(queue) > 0 {
		nums := len(queue)
		for i := 0; i < nums; i++ {
			if queue[i].Left != nil && empty != true {
				queue = append(queue, queue[i].Left)
			} else if queue[i].Left == nil && empty != true {
				empty = true
			} else if queue[i].Left != nil && empty == true {
				return false
			}
			if queue[i].Right != nil && empty != true {
				queue = append(queue, queue[i].Right)
			} else if queue[i].Right == nil && empty != true {
				empty = true
			} else if queue[i].Right != nil && empty == true {
				return false
			}
		}
		queue = queue[nums:]
	}
	return true
}