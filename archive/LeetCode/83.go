/*
83. 删除排序链表中的重复元素
2022-12-12
link: https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
question:
	给定一个已排序的链表的头 head ，
	删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
answer:
	双指针。还可以简化，待优化！
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil { // 特殊情况
		return head
	}
	low, fast := head, head
	for fast != nil {
		fast = fast.Next
		if fast.Val == low.Val { // low:指向不重复的元素，fast用来跳过重复的元素
			continue
		}
		low.Next = fast
		low = low.Next
	}
	low.Next = nil // 注意细节
	return head
}