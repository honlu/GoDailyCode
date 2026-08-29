package lcr187

import "testing"

func TestIceBreakingGame(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		m        int
		expected int
	}{
		{name: "官方示例 n=5 m=3", n: 5, m: 3, expected: 3},
		{name: "官方示例 n=10 m=17", n: 10, m: 17, expected: 2},
		{name: "只有1个人", n: 1, m: 1, expected: 0},
		{name: "只有1个人且m更大", n: 1, m: 5, expected: 0},
		{name: "每次淘汰下一个人", n: 5, m: 1, expected: 4},
		{name: "两个人 m=2", n: 2, m: 2, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := iceBreakingGame(tt.n, tt.m)
			if got != tt.expected {
				t.Errorf("iceBreakingGame(%d, %d) = %d, expected %d", tt.n, tt.m, got, tt.expected)
			}
		})
	}
}

func TestIceBreakingGameMatchesSimulation(t *testing.T) {
	for n := 1; n <= 12; n++ {
		for m := 1; m <= 15; m++ {
			got := iceBreakingGame(n, m)
			want := iceBreakingGameSimulate(n, m)
			if got != want {
				t.Fatalf("iceBreakingGame(%d, %d) = %d, simulation = %d", n, m, got, want)
			}
		}
	}
}

func iceBreakingGameSimulate(n, m int) int {
	people := make([]int, n)
	for i := 0; i < n; i++ {
		people[i] = i
	}
	idx := 0
	for len(people) > 1 {
		idx = (idx + m - 1) % len(people)
		people = append(people[:idx], people[idx+1:]...)
	}
	return people[0]
}
