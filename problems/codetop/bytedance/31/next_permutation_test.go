package bytedance31

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "ascending sequence moves to next order",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "descending sequence wraps to smallest order",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "handles duplicate values",
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
		{
			name: "changes the shortest suffix needed",
			nums: []int{1, 3, 2},
			want: []int{2, 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Fatalf("nextPermutation() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
