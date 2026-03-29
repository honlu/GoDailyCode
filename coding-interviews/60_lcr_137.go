package main

import (
	"regexp"
)

func articleMatch(s string, p string) bool {
	// 正则表达式,还有其他解法，暂时没思路
	return regexp.MustCompile("^" + p + "$").MatchString(s)
}
