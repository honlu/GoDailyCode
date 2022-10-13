/*
2、下一个更大元素1
2022-10-13
link: 496-https://leetcode.cn/problems/next-greater-element-i/
question:
	给你两个 没有重复元素 的数组 nums1 和 nums2 ，
	其中nums1 是 nums2 的子集。
	请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
answer:
	首先要定义一个res数组来存放结果，并且初始化为-1。
	其次在遍历nums2的过程中，要判断nums2[i]是否在nums1中出现过，
	因为最后是要根据nums1元素的下标来更新res数组。
	最后最关键时单调栈的使用!（类似于1.go 每日温度的代码）
*/
package main

import "fmt"

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	res := nextGreaterElement(nums1, nums2)
	fmt.Println(res)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	// 哈希存放元素-下标
	hash := make(map[int]int, len(nums1)) // key：下标元素，value:下标
	for i, v := range nums1 {
		hash[v] = i
	}
	// 单调栈,栈中存放nums2中元素对应的下标
	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		top := stack[len(stack)-1]
		for len(stack) > 0 && nums2[i] > nums2[top] {
			top = stack[len(stack)-1]
			if index, ok := hash[nums2[top]]; ok { // Go的比较特殊。查看map中是否存在这个元素，即是否是nums1中元素
				res[index] = nums2[i] // 存放下一个比它大的元素
			}
			stack = stack[:len(stack)-1] // 出栈
		}
		stack = append(stack, i) // 进栈
	}
	return res
}
