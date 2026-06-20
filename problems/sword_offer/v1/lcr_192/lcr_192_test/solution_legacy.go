//go:build legacy
// +build legacy

package v1lcr192lcr192test

import "testing"

// 写单测用例
func TestMyAtoi(t *testing.T) {
	res := myAtoi("42")
	if res != 42 {
		t.Errorf("TestMyAtoi failed, got: %d, expected: %d", res, 42)
	} else {
		t.Log("TestMyAtoi pass")
	}

}
