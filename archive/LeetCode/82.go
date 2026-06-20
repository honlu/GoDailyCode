/*
82. 删除排序链表中的重复元素2
2022-11-23
link: https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
question:
	给定一个已排序的链表的头head ， 删除原始链表中所有重复数字的节点，
	只留下不同的数字。返回 已排序的链表
answer:
	双指针+遍历+哈希表

	注意一个细节：
	某个指针变量声明之后, 还没有经过赋值时默认指向nil, 直接调用指针就会报错
	panic:runtime error: invalid memory address or nil pointer dereference
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var p, low *ListNode = new(ListNode), new(ListNode)
	hash := make(map[int]bool)
	// 注意p的初始化
	low.Next = head
	p = low           // 注意这里关键，不能设置p.Next = head
	for head != nil { // 双指针head,low
		hash[head.Val] = true
		if head.Next == nil { // 一定要注意这些细节
			break
		}
		head = head.Next
		if hash[head.Val] == true {
			for hash[head.Val] == true {
				head = head.Next
				if head == nil { // 一定要注意这些细节
					break
				}
			}
			low.Next = head
		} else {
			low = low.Next
		}
	}
	return p.Next
}