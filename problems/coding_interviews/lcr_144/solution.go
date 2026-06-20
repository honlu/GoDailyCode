package lcr144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FlipTree 翻转二叉树。
func FlipTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	line := []*TreeNode{root}
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
