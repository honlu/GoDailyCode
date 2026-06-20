package binarytree

import "container/list"

/*
4、二叉树的统一迭代法
day:2022-6-5
idea:使用栈的话，无法同时解决访问节点（遍历节点）和处理节点（将元素放进结果集）不一致的情况。
那我们就将访问的节点放入栈中，把要处理的节点也放入栈中但是要做标记。
如何标记呢，就是要处理的节点放入栈之后，紧接着放入一个空指针作为标记。 这种方法也可以叫做标记法。

type Element struct {
    // 元素保管的值
    Value interface{}
    // 内含隐藏或非导出字段
}

func (l *List) Back() *Element
*/

// 前序遍历统一迭代法
// 前序遍历：中左右， 压栈顺序：右左中
func preorderUnityTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	stack := list.New() // 初始化栈或列表
	stack.PushBack(root)
	for stack.Len() > 0 {
		element := stack.Remove(stack.Back()) // 注意这里go有一个特点当这里是interface{}类型时，后面只是使用=是不能转换为TreeNode的
		if element == nil {                   // 如果为空，则表明需要处理中间结点
			node := stack.Remove(stack.Back()).(*TreeNode) // 弹出元素，删除元素，返回删除元素的值
			res = append(res, node.Val)                    // 将中间结点加入到结果集中
			continue                                       // 继续弹出栈中下一个结点
		}
		// 不为空的话，就压栈：右左中
		node := element.(*TreeNode) // interface{}转为TreeNode指针
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		stack.PushBack(node) // 中间结点压栈后再压入nil作为中间结点的标志符
		stack.PushBack(nil)
	}
	return res
}

// 中序统一迭代遍历：左右中，压栈顺序：右中左
func inorderUnityTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		element := stack.Remove(stack.Back()) // 栈中删除栈顶的元素的Value,接口类型
		if element == nil {                   // 如果为空，则表明需要处理中间结点
			node := stack.Remove(stack.Back()).(*TreeNode) // 弹出元素，删除元素，返回删除元素的值
			res = append(res, node.Val)                    // 将中间结点加入到结果集中
			continue                                       // 继续弹出栈中下一个结点
		}
		// 不为空的话，就压栈：右中左
		node := element.(*TreeNode) // interface{}转为TreeNode指针
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		stack.PushBack(node) // 中间结点压栈后，再压入nil作为中间结点的标志符
		stack.PushBack(nil)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return res
}

// 后序统一迭代遍历：左右中， 压栈顺序：中右左
func postorderUnityTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		element := stack.Remove(stack.Back()) // 栈中删除栈顶的元素的Value,接口类型
		if element == nil {                   // 如果为空，则表明需要处理中间结点
			node := stack.Remove(stack.Back()).(*TreeNode) // 弹出元素，删除元素，返回删除元素的值
			res = append(res, node.Val)                    // 将中间结点加入到结果集中
			continue                                       // 继续弹出栈中下一个结点
		}
		// 不为空的话，就压栈：中右左
		node := element.(*TreeNode) // interface{}转为TreeNode指针
		stack.PushBack(node)        // 中间结点压栈后，再压入nil作为中间结点的标志符
		stack.PushBack(nil)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return res
}
