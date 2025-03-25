package main

/*
4-移动零
*/
func moveZeroes(nums []int)  {
	// 双指针，将前面为零和后面不为零的互换；注意非零元素相对顺序不能动
	i, j := 0, 0  // i为零所在下标，j为非零所在下标
	for i < len(nums) && nums[i] != 0{
		i++
	}
	j = i+1 // j一定在i之后
	for j < len(nums) {
		for j < len(nums) && nums[j] == 0{
			j++
		}
		for i < len(nums) && nums[i] != 0{
			i++
		}
		if i < j && j < len(nums){
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		}
	}
}
