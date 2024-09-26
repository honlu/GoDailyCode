package lcr139

/*
题目： 训练计划1
- https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

题解：
- 双指针解决
- 1. 定义两个指针，分别指向数组的头部和尾部
- 2. 头指针向后移动，直到遇到偶数
- 3. 尾指针向前移动，直到遇到奇数
- 4. 交换头尾指针指向的元素
- 5. 重复 2、3、4 步骤，直到头尾指针相遇
*/
func trainingPlan(actions []int) []int {
	p, q := 0, len(actions)-1
	for p < q {
		for p < len(actions) {
			if actions[p]%2 != 0 {
				p++
			} else {
				break
			}
		}
		for q > 0 {
			if actions[q]%2 == 0 {
				q--
			} else {
				break
			}
		}
		if p < q {
			actions[p], actions[q] = actions[q], actions[p]
		}
	}
	return actions
}

// 代码优化后的版本
func trainingPlan2(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]%2 == 1 {
			i++
		}
		for i < j && nums[j]%2 == 0 {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
