package binarytree

import "container/list"

/*
8
104、二叉树的最大深度
day :2022-6-5
update: 2023-1-16
question:
	给定一个二叉树，找出其最大深度。
	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
	说明: 叶子节点是指没有子节点的节点。
answer:
	深度和层次非常相关。所以层次遍历就可以。

	进阶：找出一条最深的路径！就比较麻烦了。
	(思路：首先获取二叉树的所有路径，以及最大深度。然后根据路径长度和深度匹配再打印。
	涉及到回溯。)
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归实现树的最大高度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 容器list模拟迭代实现-层次遍历
func maxDepthIter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()         // 保存当前层的长度，然后处理当前层。不可再for里面使用queue.Len()，防止添加元素动态变化
		for i := 0; i < size; i++ { // 依次遍历当前层在队列的头部元素
			node := queue.Remove(queue.Front()).(*TreeNode) // 头部出队列，删除元素，返回元素的值，转为TreeNode指针
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res += 1
	}
	return res
}

// 切片模拟队列，迭代实现二叉树的遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 队列初始化
	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		res++
	}
	return res
}
