package lcr131

import "testing"

func TestCuttingBamboo(t *testing.T) {
	tests := []struct {
		name      string
		bambooLen int
		expected  int
	}{
		{name: "官方示例 n=2", bambooLen: 2, expected: 1},
		{name: "长度为3", bambooLen: 3, expected: 2},
		{name: "长度为4", bambooLen: 4, expected: 4},
		{name: "长度为5", bambooLen: 5, expected: 6},
		{name: "长度为6", bambooLen: 6, expected: 9},
		{name: "长度为7", bambooLen: 7, expected: 12},
		{name: "长度为8", bambooLen: 8, expected: 18},
		{name: "长度为9", bambooLen: 9, expected: 27},
		{name: "官方示例 n=10", bambooLen: 10, expected: 36},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cuttingBamboo(tt.bambooLen)
			if got != tt.expected {
				t.Errorf("cuttingBamboo(%d) = %d, expected %d", tt.bambooLen, got, tt.expected)
			}
		})
	}
}

func TestCuttingBambooMatchesFormula(t *testing.T) {
	for n := 2; n <= 20; n++ {
		got := cuttingBamboo(n)
		want := cuttingBambooByThrees(n)
		if got != want {
			t.Fatalf("cuttingBamboo(%d) = %d, formula = %d", n, got, want)
		}
	}
}

// 贪心：尽量切成 3，余 1 时改成 2+2。
func cuttingBambooByThrees(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	prod := 1
	for n > 4 {
		prod *= 3
		n -= 3
	}
	return prod * n
}
