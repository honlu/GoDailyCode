package main

// LCR 175. 计算二叉树的深度
func calculateDepth(root *TreeNode) int {
	// 层次遍历
	var queue []*TreeNode
	var res int
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		res++
		count := len(queue)
		for i := 0; i < count; i++ {
			item := queue[i]
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
		queue = queue[count:]
	}
	return res
}
