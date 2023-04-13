/*
面试题45：把数组排成最小的数
2023-4-13 by lu
link: https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
question:
	输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
answer:
	排序+遍历
*/
func minNumber(nums []int) string {
	temp := make([]string, len(nums)) // []int转[]string
	for i, val := range nums {
		temp[i] = strconv.Itoa(val)
	}
	sort.Slice(temp, func(i, j int) bool { // 排序关键
		return temp[i]+temp[j] < temp[j]+temp[i]
	})

	return strings.Join(temp, "") // 最后结果拼接
}