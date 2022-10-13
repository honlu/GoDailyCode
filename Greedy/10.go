/*
10. 分发糖果
2022-10-13
link: 135-https://leetcode.cn/problems/candy/
question:
	老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，
	老师会根据每个孩子的表现，预先给他们评分。
	要求：每个孩子至少分配到 1 个糖果。相邻的孩子中，评分高的孩子必须获得更多的糖果。
	老师至少要准备多少颗糖果？
answer:
	采用了两次贪心的策略：
	一次是起点下标1从左到右遍历，只比较右边孩子评分比左边大的情况。只要右边比左边大，右边的糖果=左边 + 1
	一次是起点下标 length - 2从右到左遍历，只比较左边孩子评分比右边大的情况。
	只要左边比右边大，此时左边的糖果应该取本身的糖果数（符合比它左边大） 和 右边糖果数 + 1 二者的最大值，
	这样才符合它比它左边的大，也比它右边大
*/

func candy(ratings []int) int {
	res := make([]int, len(ratings))
	sum := 0
	// 第一次贪心,只考虑右边比左边是否大的情况【这是会缺少左边比右边大，但分发糖个数一样的效果】
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}
	// 第二次贪心，只考虑左边是否比右边大的情况
	for i := len(ratings) - 2; i >= 0; i-- { // 注意结束i>=0
		if ratings[i] > ratings[i+1] {
			if res[i] < res[i+1]+1 {
				res[i] = res[i+1] + 1
			}
		}
	}
	for i := 0; i < len(res); i++ {
		sum += res[i]
	}
	sum += len(ratings) // 因为每个孩子至少分发一个糖，前面初始化没有设置，所以在这里直接加上
	return sum
}