package binarytree

import (
	"sort"
)

/*
23
501、二叉搜索树中的众数
day:2022-6-10
update:2023-1-27
link：https://leetcode.cn/problems/find-mode-in-binary-search-tree/
question:
	给你一个含重复值的二叉搜索树（BST）的根节点 root ，
	找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
	如果树中有不止一个众数，可以按 任意顺序 返回。
	假定 BST 满足如下定义：
		结点左子树中所含节点的值 小于等于 当前节点的值
		结点右子树中所含节点的值 大于等于 当前节点的值
		左子树和右子树都是二叉搜索树
idea:
	如果不是二叉搜索树，最直观的方法一定是把这个树都遍历了，
	用map统计频率，把频率排个序，最后取前面高频的元素的集合。

	如果是二叉搜索树，中序遍历，然后遍历有序数组的元素出现频率，
	从头遍历，那么一定是相邻两个元素作比较，
	最后就把出现频率最高的元素输出就可以了。(注意可能有多个众数)
*/
// 普通树的递归解法
func findMode(root *TreeNode) []int {
	// map定义和声明
	hash := map[int]int{} // key：节点元素值， value:出现频率
	res := []int{}

	// 递归前序遍历申明
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		hash[root.Val]++      // 统计元素频率，中
		traversal(root.Left)  // 左
		traversal(root.Right) // 右
	}
	// 开始遍历
	traversal(root)
	// 重点，map排序
	// 拿到key
	keys := []int{}
	for k := range hash {
		keys = append(keys, k)
	}
	//  通过value对key进行从大到小的排序
	sort.Slice(keys, func(i, j int) bool {
		return hash[keys[i]] > hash[keys[j]]
	})
	// 保存众数
	res = append(res, keys[0])
	for i := 1; i < len(keys); i++ {
		if hash[keys[i]] == hash[keys[0]] {
			res = append(res, keys[i])
		} else {
			break
		}
	}
	return res
}

// 递归中序遍历，边遍历边计数，不使用额外空间。
func findMode(root *TreeNode) []int {
	var res []int
	count := 1
	max := 1
	var pre *TreeNode // 保存上一个节点
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中序遍历
		traversal(node.Left)
		if pre != nil && pre.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max { // 这里的逻辑很重要
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		pre = node
		traversal(node.Right)
	}
	traversal(root)
	return res
}
