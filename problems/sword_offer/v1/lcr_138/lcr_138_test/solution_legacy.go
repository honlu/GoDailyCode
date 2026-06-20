//go:build legacy
// +build legacy

package v1lcr138lcr138test

import "testing"

/*
单测代码模板
*/
func TestValidNumber(t *testing.T) {
	res := validNumber("0")
	if res != true {
		t.Errorf("TestValidNumber failed, got: %t, expected: %t", res, true)
	} else {
		t.Log("TestValidNumber pass")
	}
}
