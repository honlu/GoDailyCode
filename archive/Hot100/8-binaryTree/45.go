package hot100

/*
45-二叉树的右视图
https://leetcode.cn/problems/binary-tree-right-side-view/description/?envType=study-plan-v2&envId=top-100-liked
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	// 层次遍历，只要同层最右边的元素
	var res []int
	var queue []*TreeNode
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		layerNum := len(queue)
		for i := 0; i < layerNum; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[layerNum-1].Val)
		queue = queue[layerNum:]
	}
	return res
}
