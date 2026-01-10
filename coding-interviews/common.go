package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// reverseTreeNodes 反转 TreeNode 切片
// 当多个题目需要使用此功能时，统一使用此函数避免重复代码
func reverseTreeNodes(line []*TreeNode) {
	for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
		line[i], line[j] = line[j], line[i]
	}
}

// buildTree 根据层序遍历数组构建二叉树（用于测试）
// vals: 层序遍历的数组，-1表示空节点
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if i < len(vals) {
			if vals[i] != -1 { // -1 表示空节点
				node.Left = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Left)
			} else {
				queue = append(queue, nil) // 添加nil占位符
			}
			i++
		}

		// 右子节点
		if i < len(vals) {
			if vals[i] != -1 { // -1 表示空节点
				node.Right = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Right)
			} else {
				queue = append(queue, nil) // 添加nil占位符
			}
			i++
		}
	}

	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// isSameTree 比较两棵树是否相同（用于测试）
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
