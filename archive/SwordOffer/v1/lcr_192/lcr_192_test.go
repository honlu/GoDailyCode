package lcr192

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
