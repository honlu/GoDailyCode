/*
21. 多重背包理论基础
力扣上没有对应的题目，简单介绍，了解即可。

有N种物品和一个容量为V 的背包。第i种物品最多有Mi件可用，每件耗费的空间是Ci ，价值是Wi 。
求解将哪些物品装入背包可使这些物品的耗费的空间总和不超过背包容量，且价值总和最大。
多重背包和01背包是非常像的， 为什么和01背包像呢？
每件物品最多有Mi件可用，把Mi件摊开，其实就是一个01背包问题了。
熟悉在01背包的基础上写出代码！
*/

package theory

import "log"

// 多重背包可以化解为 01 背包
func multiplePack(weight, value, nums []int, bagWeight int) int {
	// 这里将多重背包改为01背包，方便后面01背包代码理解。
	for i := 0; i < len(nums); i++ {
		for nums[i] > 1 {
			weight = append(weight, weight[i])
			value = append(value, value[i])
			nums[i]--
		}
	}
	log.Println(weight)
	log.Println(value)

	res := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			res[j] = getMax(res[j], res[j-weight[i]]+value[i])
		}
		log.Println(res)
	}

	return res[bagWeight]
}

// 单元测试代码，可以看看，方便理解。
// package theory

// import "testing"

// func Test_multiplePack(t *testing.T) {
// 	type args struct {
// 		weight    []int
// 		value     []int
// 		nums      []int
// 		bagWeight int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			name: "one",
// 			args: args{
// 				weight:    []int{1, 3, 4},
// 				value:     []int{15, 20, 30},
// 				nums:      []int{2, 3, 2},
// 				bagWeight: 10,
// 			},
// 			want: 90,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := multiplePack(tt.args.weight, tt.args.value, tt.args.nums, tt.args.bagWeight); got != tt.want {
// 				t.Errorf("multiplePack() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
