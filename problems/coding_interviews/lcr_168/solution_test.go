package lcr168

import "testing"

func TestNthUglyNumber(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{name: "第1个丑数", n: 1, expected: 1},
		{name: "第2个丑数", n: 2, expected: 2},
		{name: "第3个丑数", n: 3, expected: 3},
		{name: "第4个丑数", n: 4, expected: 4},
		{name: "第5个丑数", n: 5, expected: 5},
		{name: "第6个丑数", n: 6, expected: 6},
		{name: "第7个丑数", n: 7, expected: 8},
		{name: "第8个丑数", n: 8, expected: 9},
		{name: "第9个丑数", n: 9, expected: 10},
		{name: "官方示例 n=10", n: 10, expected: 12},
		{name: "第11个丑数", n: 11, expected: 15},
		{name: "较大 n=15", n: 15, expected: 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nthUglyNumber(tt.n)
			if got != tt.expected {
				t.Errorf("nthUglyNumber(%d) = %d, expected %d", tt.n, got, tt.expected)
			}
		})
	}
}

func TestNthUglyNumberMatchesBruteForce(t *testing.T) {
	for n := 1; n <= 20; n++ {
		got := nthUglyNumber(n)
		want := nthUglyNumberBrute(n)
		if got != want {
			t.Fatalf("nthUglyNumber(%d) = %d, brute force = %d", n, got, want)
		}
	}
}

func nthUglyNumberBrute(n int) int {
	count := 0
	for x := 1; ; x++ {
		if isUgly(x) {
			count++
			if count == n {
				return x
			}
		}
	}
}

func isUgly(x int) bool {
	if x <= 0 {
		return false
	}
	for _, p := range []int{2, 3, 5} {
		for x%p == 0 {
			x /= p
		}
	}
	return x == 1
}
