package bubble_sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	got := BubbleSort([]int{3, 1, 2})
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("BubbleSort() = %v, want %v", got, want)
	}
}
