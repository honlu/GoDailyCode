package main

import "fmt"

/*
2022-9-9 360笔试第一题 ac
*/
func main() {
	// 输入
	n := 0
	fmt.Scanln(&n) // ln注意 否则会导致m[0]=0
	m := make([]int, n)
	temp := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &m[i])
		temp += m[i] // 第二种情况
	}
	// 处理 按位或操作
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			t := m[i] | m[j] // 第三种情况
			temp += t
		}
	}
	fmt.Println(temp)
}
