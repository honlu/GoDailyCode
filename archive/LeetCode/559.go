/*
559. N叉树的最大深度
2023-1-16
link: https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/
question:
	给定一个 N 叉树，找到其最大深度。
	最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
	N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
answer:
	队列，迭代，层次遍历
*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// 层次遍历模拟试试
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	// 切片模拟队列，初始化
	queue := []*Node{root}
	res := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			// 遍历n叉个节点
			for j := 0; j < len(node.Children); j++ {
				if node.Children[j] != nil {
					queue = append(queue, node.Children[j])
				}
			}
		}
		// 层次更改
		res++
		queue = queue[size:]
	}
	return res
}