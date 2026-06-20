/*
206. 反转链表
2023-1-6
link: https://leetcode.cn/problems/reverse-linked-list/
question:
	给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
answer:
	为了保证头节点一致，设置两个指针：一个虚拟头节点和当前节点来完成。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法一：双指针，不够优化
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil{
//         return head
//     }
// 	var p *ListNode
// 	cur := head
// 	for cur.Next != nil{
// 		temp := cur.Next
// 		cur.Next = p
// 		p = cur
// 		cur = temp
// 	}
// 	cur.Next = p
// 	return cur
// }

// 解法二：上面代码再优化
// func reverseList(head *ListNode) *ListNode {
// 	var p *ListNode // 注意这里p为nil，不能初始化、创建内存
// 	cur := head
// 	for cur != nil { // 注意这里优化
// 		temp := cur.Next
// 		cur.Next = p
// 		p = cur
// 		cur = temp
// 	}
// 	return p // 注意这里返回值
// }

// 解法三：再进一步提高，使用递归方法.当cur为空的时候循环结束，不断将cur指向pre的过程。
// 首先定义一个递归方法，可以在函数内，也可以单独
func Recursion(pre, cur *ListNode) *ListNode {
	if cur == nil { // 递归结束条件
		return pre
	}
	temp := cur.Next
	cur.Next = pre
	// 可以和双指针法的代码进行对比，如下递归的写法，其实就是做了这两步
	// pre = cur, cur = temp
	return Recursion(cur, temp)
}

// 关键是初始化的地方，可能会不理解， 可以看到双指针法中初始化 cur = head，pre = NULL，
// 在递归法中可以从如下代码看出初始化的逻辑也是一样的，只不过写法变了。
func reverseList(head *ListNode) *ListNode {
	return Recursion(nil, head) // 和双指针初始化一样的逻辑，pre=nil,cur=head
}