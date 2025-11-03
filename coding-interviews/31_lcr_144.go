package main

// flipTree LCR 144 - 翻转二叉树
func flipTree(root *TreeNode) *TreeNode {
	// 层次遍历，替换左右节点
	if root == nil {
		return root
	}
	var line []*TreeNode
	line = append(line, root)
	for len(line) > 0 {
		count := len(line)
		for i := 0; i < count; i++ {
			item := line[i]
			item.Left, item.Right = item.Right, item.Left
			if item.Left != nil {
				line = append(line, item.Left)
			}
			if item.Right != nil {
				line = append(line, item.Right)
			}
		}
		line = line[count:]
	}
	return root
}
