/*
199.二叉树的右视图
2022-11-21
link:https://leetcode.cn/problems/binary-tree-right-side-view/
question:
	给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，
	返回从右侧所能看到的节点值。
answer:
	注意：关键在于理解题目右侧底部含义，右侧是指最右侧的节点。还有最底部右侧的节点呢？
	代码测试一下! 个人理解是指最右侧的节点（即使不是最底部！）
	然而测试代码错误！

	所以本题是要打印出每层最后一个节点的值！
	方法：层次遍历保存每一层的节点值，再依次打印或保存每层最后一个节点值！
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 简单版本：只输出最右侧的节点（非通过版本！）
// func rightSideView(root *TreeNode) []int {
// 	var res []int
// 	if root == nil {
// 		return res
// 	}
// 	// 适用队列保存节点
// 	var queue []*TreeNode
// 	queue = append(queue, root)
// 	for len(queue) != 0 {
// 		res = append(res, queue[0].Val)
// 		if queue[0].Right != nil {
// 			queue[0] = queue[0].Right
// 		} else {
// 			break
// 		}
// 	}
// 	return res
// }
// 加强版。层次遍历，然后将每层最后一个节点提取出来汇总到结果中！
func rightSideView(root *TreeNode) []int {
	var res []int
	var level [][]int
	if root == nil {
		return res
	}
	// 适用队列保存每层最右侧节点
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		s := len(queue)
		temp := []int{}
		// 保存当前层的值和下一层的节点
		for i := 0; i < s; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		level = append(level, temp)
		// 删除当前层的节点
		queue = queue[s:]
	}
	// 打印或保存每层最后一个节点值！
	for i := 0; i < len(level); i++ {
		res = append(res, level[i][len(level[i])-1])
	}
	return res
}