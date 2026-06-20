package lcr150

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder 返回二叉树的层序遍历结果。
func LevelOrder(root *TreeNode) [][]int {
	var line []*TreeNode
	var res [][]int
	if root == nil {
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
