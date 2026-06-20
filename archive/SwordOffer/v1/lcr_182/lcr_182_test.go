package lcr182

import "testing"

func TestDynamicPassword(t *testing.T) {
	res := dynamicPassword("s3cur1tyC0d3", 4)
	if res != "r1tyC0d3s3cu" {
		t.Errorf("TestDynamicPassword failed, got: %s, expected: %s", res, "r1tyC0d3s3cu")
	} else {
		t.Log("TestDynamicPassword pass")
	}
}
