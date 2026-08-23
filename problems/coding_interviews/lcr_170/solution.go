package lcr170

// 交易逆序对的总数
// 双层for循环暴力超时
// 优化：分治法
// 将数组分成两部分，分别计算两部分的逆序对数，然后计算跨越两部分的逆序对数
// 跨越两部分的逆序对数可以通过归并排序来计算
// 归并排序的时间复杂度是O(nlogn)
func reversePairs(record []int) int {
	// 采用归并排序的思想，归并排序中要置换的次数就是答案
	var merge func(record []int, start, end int) int
	merge = func(record []int, start, end int) int {
		if start >= end {
			return 0
		}
		mid := start + (end-start)/2
		cnt := merge(record, start, mid) + merge(record, mid+1, end)
		temp := []int{}
		i, j := start, mid+1
		for i <= mid && j <= end {
			if record[i] <= record[j] {
				temp = append(temp, record[i])
				i++
			} else {
				cnt += mid - i + 1 // 假设左右两边已经给你有序，由于左边已经有序
				temp = append(temp, record[j])
				j++
			}
		}
		if i <= mid { // 注意等于号
			temp = append(temp, record[i:mid+1]...)
		}
		if j <= end {
			temp = append(temp, record[j:end+1]...)
		}
		copy(record[start:end+1], temp) // 注意这里需要修改原来的record数组
		return cnt
	}
	return merge(record, 0, len(record)-1) // 注意使用闭包范围
}
