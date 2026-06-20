package hot100

/*
44-二叉搜索树中第k小的元素
https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/?envType=study-plan-v2&envId=top-100-liked
*/
func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历保存结果，当结果队列长度等k返回元素
	var res []int
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		preOrder(node.Left)
		res = append(res, node.Val)
		if len(res) == k {
			return
		}
		preOrder(node.Right)
	}
	preOrder(root)
	return res[k-1]
}
