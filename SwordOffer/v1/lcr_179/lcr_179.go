package lcr179

/*
题目：查找总价格为目标值的两个数
- https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
- 输入一个递增排序的数组和一个数字 s，在数组中查找两个数，使得它们的和正好是 s。如果有多对数字的和等于 s，则输出任意一对即可。

题解：
- 双指针解决
- 1. 定义两个指针，分别指向数组的头部和尾部
- 2. 如果两个指针指向的元素之和等于目标值，则返回两个元素
- 3. 如果两个指针指向的元素之和小于目标值，则头指针向后移动
- 4. 如果两个指针指向的元素之和大于目标值，则尾指针向前移动
- 5. 重复 2、3、4 步骤，直到头尾指针相遇
*/
func twoSum(price []int, target int) []int {
	i, j := 0, len(price)-1
	var res []int
	for i < j {
		if price[i]+price[j] == target {
			res = append(res, []int{price[i], price[j]}...)
			break
		} else if price[i]+price[j] > target {
			j--
		} else {
			i++
		}
	}

	return res
}
