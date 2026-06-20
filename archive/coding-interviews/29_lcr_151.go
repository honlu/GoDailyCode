package main

// zigzagLevelOrder LCR 151 - 之字形层序遍历（从上到下打印二叉树III）
func zigzagLevelOrder(root *TreeNode) [][]int {
	i := 0
	var line []*TreeNode
	var res [][]int
	if root == nil {
		return res
	}
	line = append(line, root)
	for len(line) > 0 {
		i++
		var tempArr []int
		count := len(line)
		for item := 0; item < count; item++ {
			temp := line[item]
			tempArr = append(tempArr, temp.Val)
			if i%2 == 1 { // 奇数层，它的左右字数顺序入
				if temp.Left != nil {
					line = append(line, temp.Left)
				}
				if temp.Right != nil {
					line = append(line, temp.Right)
				}
			} else { // 偶数层，它的左右子树逆序入
				if temp.Right != nil {
					line = append(line, temp.Right)
				}
				if temp.Left != nil {
					line = append(line, temp.Left)
				}
			}
		}
		line = line[count:]
		reverseTreeNodes(line) // 反转当前层节点，使用 common.go 中的公共函数
		res = append(res, tempArr)
	}
	return res
}
