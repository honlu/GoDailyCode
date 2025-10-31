package main

// levelOrder LCR 150 - 层序遍历（从上到下打印二叉树）
func levelOrder(root *TreeNode) [][]int {
	// 层次遍历
	var line []*TreeNode
	var res [][]int
	if root == nil { // 边界条件
		return res
	}
	line = append(line, root)
	for len(line) > 0 {
		var temp []int
		num := len(line)
		for i := 0; i < num; i++ {
			temp = append(temp, line[i].Val)
			if line[i].Left != nil {
				line = append(line, line[i].Left)
			}
			if line[i].Right != nil {
				line = append(line, line[i].Right)
			}
		}
		res = append(res, temp)
		line = line[num:]
	}
	return res
}
