package lcr158

import "testing"

func TestInventoryManagement(t *testing.T) {
	tests := []struct {
		name  string
		stock []int
		want  int
	}{
		{name: "majority in middle after sort", stock: []int{6, 1, 3, 1, 1, 1}, want: 1},
		{name: "single item", stock: []int{7}, want: 7},
		{name: "majority at high value", stock: []int{2, 2, 1, 2, 3}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InventoryManagement(tt.stock)
			if got != tt.want {
				t.Fatalf("InventoryManagement() = %d, want %d", got, tt.want)
			}
		})
	}
}
