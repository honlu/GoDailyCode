/*
113.路径总和2
2023-1-23
link:https://leetcode.cn/problems/path-sum-ii/
question:
	给你二叉树的根节点 root 和一个整数目标和 targetSum ，
	找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
	叶子节点是指没有子节点的节点。

answer:
	(难点：从是否存在，到打印多个符合要求的路径)
	要遍历整个树，找到所有路径，所以递归函数不要返回值！
	层次遍历：使用三个个队列，分别存储将要遍历的节点，
		，存储根节点到这些节点的路径和即，
		以及路径。
*/
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	// 三个队列初始化
	queueNode := []*TreeNode{root} // 要遍历的节点
	queueVal := []int{root.Val}    // 根节点到当前节点的总和
	queuePath := [][]int{}         // 根节点到当前节点的路径
	queuePath = append(queuePath, []int{root.Val})
	// 层次遍历开始
	for len(queueNode) > 0 {
		size := len(queueNode)
		for i := 0; i < size; i++ {
			node, val := queueNode[i], queueVal[i]
			path := queuePath[i]
			if node.Left == nil && node.Right == nil { // 叶子节点
				if val == targetSum { // 判断根节点到当前叶子节点的路径和是否满足要求
					// 将满足此条件的路径添加到res中
					res = append(res, path)
				}
			}
			if node.Left != nil {
				queueNode = append(queueNode, node.Left)
				queueVal = append(queueVal, val+node.Left.Val)
				// 注意：路径添加
				temp := append([]int{}, path...)
				temp = append(temp, node.Left.Val) // 追加多个元素
				queuePath = append(queuePath, temp)
			}
			if node.Right != nil {
				queueNode = append(queueNode, node.Right)
				queueVal = append(queueVal, val+node.Right.Val)
				temp := append([]int{}, path...) // 追加多个元素
				temp = append(temp, node.Right.Val)
				queuePath = append(queuePath, temp)
			}
		}
		queueNode = queueNode[size:]
		queueVal = queueVal[size:]
		queuePath = queuePath[size:]
	}
	return res
}