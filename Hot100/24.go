package hot100

/*
24-回文链表
https://leetcode.cn/problems/palindrome-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
*/
func isPalindrome(head *ListNode) bool {
	var nums []int
	p := head
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}
