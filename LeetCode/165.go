/*
165. 比较版本号
2022-11-30
link:
	https://leetcode.cn/problems/compare-version-numbers/
question:
	给你两个版本号 version1 和 version2 字符串形式，请你比较它们。
answer:
	字符串切割，字符串数字比较
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")
	res := 0
	i := 0
	for i < len(s1) && i < len(s2) { // 比较相同长度部分的修订号
		t1, _ := strconv.Atoi(s1[i])
		t2, _ := strconv.Atoi(s2[i])
		i++
		if t1 > t2 {
			res = 1
			return res
		} else if t1 < t2 {
			res = -1
			return res
		} else {
			continue
		}
	}
	// 当有修订号个数前面都一样时，要比较后面是否有大于0的情况
	for i < len(s1) {
		t1, _ := strconv.Atoi(s1[i])
		i++
		if t1 > 0 {
			res = 1
			return res
		}
	}
	for i < len(s2) {
		t2, _ := strconv.Atoi(s2[i])
		i++
		if t2 > 0 {
			res = -1
			return res
		}
	}
	return res
}

func main() {
	version1 := "1.01"
	version2 := "1.001"
	res := compareVersion(version1, version2)
	fmt.Println(res)
}
