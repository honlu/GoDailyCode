/*
8
93. 复原IP地址
day:2022-10-8
update: 2023-2-4 by lu
link: https://leetcode.cn/problems/restore-ip-addresses/
question：
	给定一个只包含数字的字符串，
	复原它并返回所有可能的 IP 地址格式。
answer:
	有点类似于分割字符串。
	而切割问题就可以使用回溯搜索法把所有可能性搜出来。
*/
package main

import (
	"strconv"
	"strings"
)

// 一个函数完成-回溯
func restoreIpAddresses(s string) []string {
	var res []string  // 符合的结果集
	var path []string // 符合的结果，用[]string来存储符合的每段整数,方便添加和回溯，以及最后拼接成ip地址
	var backTrack func(start int)
	backTrack = func(start int) {
		// base case
		if len(path) == 4 { // path正好每个元素都符合0-255的条件
			if start == len(s) {
				str := strings.Join(path, ".") // 字符串拼接
				res = append(res, str)
				return
			}
		}
		// 回溯标准，递归逻辑
		for i := start; i < len(s); i++ {
			if s[start] == '0' && i != start {
				break // 含有前导0，无效
			}
			// 处理
			str := s[start : i+1]
			num, _ := strconv.Atoi(str)
			if num >= 0 && num <= 255 {
				path = append(path, str) // 符合条件，就进入下一层
				backTrack(i + 1)
				path = path[:len(path)-1] // 回溯
			} else { // 如果不满足条件，往后也不可能满足条件，直接退出
				break
			}
		}
	}
	backTrack(0)

	return res
}

// 拆开写
func restoreIpAddresses(s string) []string {
	var res, track []string
	backTrack(0, s, track, &res)
	return res
}

// 递归函数-回溯
func backTrack(startIndex int, s string, track []string, res *[]string) {
	// 递归的终止条件
	if startIndex == len(s) && len(track) == 4 {
		temp := track[0] + "." + track[1] + "." + track[2] + "." + track[3] // 一个正常IP
		*res = append(*res, temp)
	}
	// for 循环
	for i := startIndex; i < len(s); i++ {
		// 处理-切片插入
		track := append(track, s[startIndex:i+1])
		if i-startIndex+1 <= 3 && len(track) <= 4 && isNormalIp(startIndex, i, s) {
			// 递归
			backTrack(i+1, s, track, res)
		} else { // 如果首尾超过3个，或路径多余4个，或前导为0
			return
		}
		// 回溯
		track = track[:len(track)-1]
	}
}

// 判断是否是正常的IP整数
func isNormalIp(startIndex, end int, s string) bool {
	checkInt, _ := strconv.Atoi(s[startIndex : end+1]) // 字符串转整数 string to int
	if end-startIndex+1 > 1 && s[startIndex] == '0' {  // 对于前导0的IP（特别注意s[startIndex]=='0'的判断，不能写成数字0的判断）
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}
