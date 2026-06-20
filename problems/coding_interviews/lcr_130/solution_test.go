package lcr130

import "testing"

func TestWardrobeFinishing(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		cnt  int
		want int
	}{
		{name: "small grid all reachable", m: 3, n: 3, cnt: 5, want: 9},
		{name: "limited by digit sum", m: 3, n: 3, cnt: 1, want: 3},
		{name: "single cell", m: 1, n: 1, cnt: 0, want: 1},
		{name: "larger grid", m: 10, n: 10, cnt: 5, want: 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WardrobeFinishing(tt.m, tt.n, tt.cnt); got != tt.want {
				t.Fatalf("WardrobeFinishing() = %d, want %d", got, tt.want)
			}
		})
	}
}
