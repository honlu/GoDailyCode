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

import "strconv"

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
