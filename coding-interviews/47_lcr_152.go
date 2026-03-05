package main

func verifyTreeOrder(postorder []int) bool {
	// 数组一定是二叉树的后序遍历，主要看这个二叉树是不是搜索树，即左子树都比中间节点小，右子树都比中间节点大
	// 递归实现
	m := len(postorder)
	if m == 0 {
		return true
	}
	root := postorder[m-1]
	i := 0
	for i < m-1 {
		if postorder[i] < root {
			i++
		} else if postorder[i] == root {
			return false
		} else {
			break
		}
	}
	left := postorder[:i]
	right := postorder[i : m-1]
	// 如果右子树有比根小，或左子树有比根大的，返回false
	for _, item := range left {
		if item >= root {
			return false
		}
	}
	for _, item := range right {
		if item <= root {
			return false
		}
	}
	return verifyTreeOrder(left) && verifyTreeOrder(right)
}
