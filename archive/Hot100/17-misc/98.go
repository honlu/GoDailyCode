package misc

/*
98-75.颜色分类
https://leetcode.cn/problems/sort-colors/?envType=study-plan-v2&envId=top-100-liked
*/
/*
思路：原地排序算法
*/
func sortColors(nums []int) {
	// 手动实现原地排序算法，冒泡排序
	for i := 0; i < len(nums); i++ { // 第i轮，找到倒数第i个最大值
		for j := 0; j < len(nums)-i-1; j++ { //
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
