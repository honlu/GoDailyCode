package binarytree

/*
19
617、合并二叉树
day:2022-6-9
update:2023-1-25
link:https://leetcode.cn/problems/merge-two-binary-trees/
question:
	给你两棵二叉树： root1 和 root2 。
	想象一下，当你将其中一棵覆盖到另一棵之上时，
	两棵树上的一些节点将会重叠（而另一些不会）。
	你需要将这两棵树合并成一棵新二叉树。合并的规则是：
	如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；
	否则，不为 null 的节点将直接作为新二叉树的节点。
	返回合并后的二叉树。

idea:
	合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，
	否则不为 NULL 的节点将直接作为新二叉树的节点。
	其实和遍历一个树逻辑是一样的，只不过传入两个树的节点，同时操作。
	哪种遍历都是可以的。
*/

// 前序遍历-简洁版，以第一个二叉树为改动
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 递归结束条件，在root1上做改动
	if root1 == nil {
		return root2 // 注意这里返回的节点
	}
	if root2 == nil {
		return root1
	}
	// 单层递归的逻辑
	root1.Val += root2.Val                             // 中
	root1.Left = mergeTrees(root1.Left, root2.Left)    // 左
	root1.Right = mergeTrees(root1.Right, root2.Right) // 右

	return root1
}

// 迭代版的话，需要队列或栈来存储当前节点(模拟层次遍历实现)
// 就是把两个树的节点同时加入队列进行比较以及合并
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 队列初始化
	queue := []*TreeNode{root1, root2}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i += 2 {
			node1, node2 := queue[i], queue[i+1]
			// 此时两个节点一定不为空，val相加
			node1.Val += node2.Val

			// 如果两棵树左节点都不为空，加入队列
			if node1.Left != nil && node2.Left != nil {
				queue = append(queue, node1.Left)
				queue = append(queue, node2.Left)
			}
			// 如果两颗树右节点都不为空，加入队列
			if node1.Right != nil && node2.Right != nil {
				queue = append(queue, node1.Right)
				queue = append(queue, node2.Right)
			}
			// 当node1的左节点为空，node2的左节点不为空，就赋值过去
			if node1.Left == nil && node2.Left != nil {
				node1.Left = node2.Left // 注意这里就不加队列了！
			}
			// 当node1的右节点为空，node2的右节点不为空，就赋值过去
			if node1.Right == nil && node2.Right != nil {
				node1.Right = node2.Right
			}
		}
		queue = queue[size:]
	}

	return root1
}
