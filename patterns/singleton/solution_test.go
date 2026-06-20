package singleton

import "testing"

func TestSingletonInstances(t *testing.T) {
	if GetInstance() != GetInstance() {
		t.Fatal("GetInstance should return the same instance")
	}
	if GetLazyInstance() != GetLazyInstance() {
		t.Fatal("GetLazyInstance should return the same instance")
	}
}
