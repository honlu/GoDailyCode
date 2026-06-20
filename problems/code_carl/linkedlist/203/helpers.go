package linkedlist203

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val      int
	Next     *Node
	Random   *Node
	Children []*Node
	Left     *Node
	Right    *Node
}
