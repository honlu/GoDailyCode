package main

import (
	"sort"
	"strconv"
	"strings"
)

func crackPassword(password []int) string {
	// 数字的排序方法：0最小
	temp := make([]string, 0, len(password)) // 使用切片而不是数组，因为数组长度是固定的，而切片长度是动态的
	for _, item := range password {
		temp = append(temp, strconv.Itoa(item))
	}
	sort.Slice(
		temp, func(i, j int) bool {
			return temp[i]+temp[j] < temp[j]+temp[i]
		},
	)
	return strings.Join(temp, "")
}
