package hot100

/*
48-路径总和3
https://leetcode.cn/problems/path-sum-iii/description/?envType=study-plan-v2&envId=top-100-liked
*/
/*
层次遍历：
1. 层次遍历原先只记录各层节点。
2. 为解本次需求还需要记录从根节点到各个当前节点的路径，然后根据路径逆向计算和和目标值比较
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	// 层次遍历
	if root == nil {
		return 0
	}
	var res int
	queueNode := []*TreeNode{root} // 层次遍历节点
	queuePath := [][]int{}         // 根节点到当前节点的路径
	queuePath = append(queuePath, []int{root.Val})
	for len(queueNode) > 0 {
		size := len(queueNode)
		for i := 0; i < size; i++ {
			node := queueNode[i]
			path := queuePath[i]
			// 逆推路径和是否存在
			sum := 0
			for j := len(path) - 1; j >= 0; j-- {
				sum += path[j]
				if sum == targetSum {
					res++
				}
			}
			// 追加新的节点和节点路径
			if node.Left != nil {
				queueNode = append(queueNode, node.Left)
				temp := append([]int{}, path...) // 避免修改path导致底层slice问题
				temp = append(temp, node.Left.Val)
				queuePath = append(queuePath, temp)
			}
			if node.Right != nil {
				queueNode = append(queueNode, node.Right)
				temp := append([]int{}, path...)
				temp = append(temp, node.Right.Val)
				queuePath = append(queuePath, temp)
			}
		}
		queueNode = queueNode[size:] // 回撤
		queuePath = queuePath[size:]
	}
	return res
}
