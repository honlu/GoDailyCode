package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 本题主要考察二叉树的遍历，然后序列化；再然后如何根据序列化的内容怎么反序列化成二叉树
	var queue []string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			queue = append(queue, "nil")
			return
		}
		queue = append(queue, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(queue, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	var index int
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if index >= len(list) {
			return nil
		}
		val := list[index]
		index++
		if val == "nil" {
			return nil
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil
		}
		return &TreeNode{Val: num, Left: dfs(), Right: dfs()}
	}
	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
