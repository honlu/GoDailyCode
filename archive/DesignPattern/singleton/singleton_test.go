package singleton

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// 饿汉式的单元测试
func TestGetInstance(t *testing.T) { // go test -run TestGetInstance 进行测试运行
	assert.Equal(t, GetInstance(), GetInstance(), "Add Error!") // 判断两个实例是否一样。因为单例模式，一个类只能创建一个对象（或者叫实例）
}

func BenchmarkGetInstanceParallel(b *testing.B) { // 基准测试 go
	b.RunParallel(func(p *testing.PB) { // RunParallel 测试并发性能
		for p.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
