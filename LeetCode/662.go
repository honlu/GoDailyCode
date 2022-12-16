/*
662. 二叉树最大宽度
2022-12-14
link: https://leetcode.cn/problems/maximum-width-of-binary-tree/
question:
	给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
	树的 最大宽度 是所有层中最大的 宽度 。

	每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。
	将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的 null 节点，
	这些 null 节点也计入长度。
answer:
	难点：空节点算入。
	广度优先搜索（层次遍历），由于算入空节点，所以可以看作完全二叉树。
	利用完全二叉树的特点，父节点的编号为i,其两个孩子节点分别为2i和2i+1。
	每层宽度 = 用每层节点的最大编号 - 最小编号+1
	参考：https://leetcode.cn/problems/maximum-width-of-binary-tree/solution/er-cha-shu-zui-da-kuan-du-by-leetcode-so-9zp3/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type pair struct { // 关键：组件信息的结构体
	node  *TreeNode
	index int // 编号
}

func widthOfBinaryTree(root *TreeNode) int {
	var queue []pair
	queue = append(queue, pair{root, 1})
	res := 1
	for queue != nil {
		if res < queue[len(queue)-1].index-queue[0].index+1 { // 和当前层宽度进行比较
			res = queue[len(queue)-1].index - queue[0].index + 1
		}
		temp := queue            // 当前层
		queue = nil              // 保存下一层
		for _, t := range temp { // 遍历当前层,添加下一层节点
			if t.node.Left != nil {
				queue = append(queue, pair{t.node.Left, t.index * 2})
			}
			if t.node.Right != nil {
				queue = append(queue, pair{t.node.Right, t.index*2 + 1})
			}
		}
	}
	return res
}