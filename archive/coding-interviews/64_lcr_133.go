package main

import "fmt"

func hammingWeight(num uint32) int {
	// 将无符号整数转为二进制字符串
	x := fmt.Sprintf("%b", num)
	var res int
	for _, item := range x {
		if item == '1' {
			res++
		}
	}
	return res
}
