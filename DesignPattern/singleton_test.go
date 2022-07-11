package singleton

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// 饿汉式的单元测试
func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInstance()) //?
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if GetInstance() != GetInstance() { // ?
				b.Errorf("test fail")
			}
		}
	})
}
