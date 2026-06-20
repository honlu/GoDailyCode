package lcr129

import "testing"

func TestWordPuzzle(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]byte
		target string
		want   bool
	}{
		{name: "sample path exists", grid: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, target: "ABCCED", want: true},
		{name: "vertical path exists", grid: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, target: "SEE", want: true},
		{name: "cannot reuse cell", grid: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, target: "ABCB", want: false},
		{name: "single cell", grid: [][]byte{{'A'}}, target: "A", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordPuzzle(tt.grid, tt.target); got != tt.want {
				t.Fatalf("WordPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}
