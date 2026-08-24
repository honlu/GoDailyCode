package lcr160

import (
	"math"
	"testing"
)

const medianTolerance = 1e-9

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < medianTolerance
}

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name  string
		steps []medianStep
	}{
		{
			name: "官方示例",
			steps: []medianStep{
				{add: 1},
				{add: 2, find: 1.5, check: true},
				{add: 3, find: 2.0, check: true},
			},
		},
		{
			name: "单元素",
			steps: []medianStep{
				{add: 7, find: 7.0, check: true},
			},
		},
		{
			name: "升序数据流",
			steps: []medianStep{
				{add: 1, find: 1.0, check: true},
				{add: 2, find: 1.5, check: true},
				{add: 3, find: 2.0, check: true},
				{add: 4, find: 2.5, check: true},
			},
		},
		{
			name: "降序数据流",
			steps: []medianStep{
				{add: 5, find: 5.0, check: true},
				{add: 4, find: 4.5, check: true},
				{add: 3, find: 4.0, check: true},
			},
		},
		{
			name: "含负数和重复值",
			steps: []medianStep{
				{add: -1},
				{add: -1, find: -1.0, check: true},
				{add: 0, find: -1.0, check: true},
				{add: 1, find: -0.5, check: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := Constructor()
			for i, step := range tt.steps {
				mf.AddNum(step.add)
				if !step.check {
					continue
				}
				got := mf.FindMedian()
				if !floatEqual(got, step.find) {
					t.Fatalf("step %d after AddNum(%d): FindMedian() = %v, want %v", i, step.add, got, step.find)
				}
			}
		})
	}
}

type medianStep struct {
	add   int
	find  float64
	check bool
}
