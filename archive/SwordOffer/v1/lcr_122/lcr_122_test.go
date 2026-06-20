package lcr_122

import "testing"

func TestPathEncrypting(t *testing.T) {
	res := pathEncrypting("a.aef.qerf.bb")
	if res != "a aef qerf bb" {
		t.Errorf("TestPathEncrypting failed, got: %s, expected: %s", res, "a aef qerf bb")
	} else {
		t.Log("TestPathEncrypting pass")
	}
}
