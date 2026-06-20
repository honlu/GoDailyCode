/*
468. 验证IP地址
2022-12-19
link: https://leetcode.cn/problems/validate-ip-address/
question:
	给定一个字符串 queryIP。如果是有效的 IPv4 地址，返回 "IPv4" ；
	如果是有效的 IPv6 地址，返回 "IPv6" ；
	如果不是上述类型的 IP 地址，返回 "Neither" 。

answer:
	字符串处理问题:分割，包含，比较。
*/
package leetcode

import (
	"strings"
)

func validIPAddress(queryIP string) string {
	if strings.Contains(queryIP, ".") && ipv4(queryIP) {
		return "IPV4"
	} else if strings.Contains(queryIP, ":") && ipv6(queryIP) {
		return "IPV6"
	} else {
		return "Neither"
	}
}

func ipv4(queryIP string) bool {
	s4 := strings.Split(queryIP, ".")
	if len(s4) != 4 {
		return false
	}
	// 不能包含前导零，怎么处理呢？
	for _, val := range s4 {
		if len(val) > 3 { // 小心有个特例，贼长的1"11111111111111...11111111111111111"
			return false
		}
		if val >= "0" && val <= "255" {
			if val == "0" && len(val) > 1 {
				return false
			} else if val > "0" && val[0] == '0' {
				return false
			}
			continue
		} else {
			return false
		}
	}
	return true
}

func ipv6(queryIP string) bool {
	s6 := strings.Split(queryIP, ":")
	if len(s6) != 8 {
		return false
	}
	// 可以包含数字，小写和大写a-f的字母
	for _, val := range s6 {
		if len(val) >= 1 && len(val) <= 4 {
			for _, i := range val {
				if (i >= '0' && i <= '9') || (i >= 'a' && i <= 'f') || (i >= 'A' && i <= 'F') {
					continue
				} else {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true
}
