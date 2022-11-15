/*
103. 二叉树的锯齿形层序遍历
2022-11-15
link: https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
question:
	给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。
	（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
answer:
	层次遍历的变形！
	迭代实现！队列适合模拟广度优先遍历，栈适合模拟深度优先遍历！
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	// 队列创建
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	flag := 0 // 偶数层0：层从左到右， 奇数层1：层从右到左
	// 遍历
	for len(queue) != 0 {
		size := len(queue) // 当前层节点数
		var temp []int
		// 依次遍历
		for i := 0; i < size; i++ {
			node := queue[i]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		if flag%2 == 1 {
			// 反转本层节点，再加入到res中
			for j, k := 0, len(temp)-1; j < k; j, k = j+1, k-1 {
				temp[j], temp[k] = temp[k], temp[j]
			}
		}
		res = append(res, temp)
		flag += 1 // 方便判断层序
	}
	return res
}