package lcr159

import (
	"reflect"
	"testing"
)

func TestInventoryManagement(t *testing.T) {
	tests := []struct {
		name  string
		stock []int
		cnt   int
		want  []int
	}{
		{name: "take two smallest", stock: []int{4, 1, 6, 2, 5}, cnt: 2, want: []int{1, 2}},
		{name: "take none", stock: []int{3, 2, 1}, cnt: 0, want: []int{}},
		{name: "take all sorted", stock: []int{2, 1}, cnt: 2, want: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InventoryManagement(tt.stock, tt.cnt)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("InventoryManagement() = %v, want %v", got, tt.want)
			}
		})
	}
}
