package heap

/*
76-数据流的中位数
https://leetcode.cn/problems/find-median-from-data-stream/?envType=study-plan-v2&envId=top-100-liked

*/

// 以下简单添加，排序。超时。
// type MedianFinder struct {
// 	queue []int
// }

// func Constructor() MedianFinder {
// 	return MedianFinder{}
// }

// func (this *MedianFinder) AddNum(num int) {
// 	this.queue = append(this.queue, num)
// 	sort.Ints(this.queue)

// }

// func (this *MedianFinder) FindMedian() float64 {
// 	if len(this.queue)%2 == 0 {
// 		return float64(this.queue[(len(this.queue)/2)-1]+this.queue[len(this.queue)/2]) * 1.0 / 2.0
// 	} else {
// 		return float64(this.queue[len(this.queue)/2])
// 	}
// }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
