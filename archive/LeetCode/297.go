/*
297. 二叉树的序列化与反序列化
2022-12-17
link:
question:
	序列化是将一个数据结构或者对象转换为连续的比特位的操作，
	进而可以将转换后的数据存储在一个文件或者内存中，
	同时也可以通过网络传输到另一个计算机环境，
	采取相反方式重构得到原数据。

	请设计一个算法来实现二叉树的序列化与反序列化。
	这里不限定你的序列 / 反序列化算法执行逻辑,你只需要保证一个
	二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
answer:
	二叉树遍历和搜索相关
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct{}

func Constructor() Codec {
	return Codec{} // 返回一个
}

// Serializes a tree to a single string.
// 先序遍历+递归
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	return strings.Join([]string{strconv.Itoa(root.Val), this.serialize(root.Left),
		this.serialize(root.Right)}, ",") // 注意这里序列化的形式
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var pre = strings.Split(data, ",")
	var root *TreeNode // 根节点，返回值
	var cur = &root    // cur 是**TreeNode类型
	var stack = []**TreeNode{cur}
	for _, val := range pre {
		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		num, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		temp := &TreeNode{Val: num}
		*cur, stack = temp, append(stack, &temp.Right, &temp.Left) // 注意顺序
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */