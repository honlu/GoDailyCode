package binarytree

import "container/list"

/*
5、二叉树的层序遍历
day:2022-6-5
idea:使用队列实现二叉树广度优先遍历
借用一个辅助数据结构即队列来实现，队列先进先出，符合一层一层遍历的逻辑，而是用栈先进后出适合模拟深度优先遍历也就是递归的逻辑。
而这种层序遍历方式就是图论中的广度优先遍历，只不过我们应用在二叉树上。

队列：先进先出
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}    // 结果切片
	queue := list.New() // 队列
	queue.PushBack(root)
	var tempArr []int // 层切片
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
			tempArr = append(tempArr, node.Val)
		}
		res = append(res, tempArr) // 加入结果集
		tempArr = []int{}          // 清楚层的数据
	}
	return res
}
