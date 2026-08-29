package lcr162

import "testing"

func TestDigitOneInNumber(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected int
	}{
		{name: "n=0", num: 0, expected: 0},
		{name: "n=1", num: 1, expected: 1},
		{name: "n=10", num: 10, expected: 2},
		{name: "n=11", num: 11, expected: 4},
		{name: "官方示例 n=12", num: 12, expected: 5},
		{name: "官方示例 n=13", num: 13, expected: 6},
		{name: "n=20", num: 20, expected: 12},
		{name: "n=99", num: 99, expected: 20},
		{name: "n=100", num: 100, expected: 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := digitOneInNumber(tt.num)
			if got != tt.expected {
				t.Errorf("digitOneInNumber(%d) = %d, expected %d", tt.num, got, tt.expected)
			}
		})
	}
}

func TestDigitOneInNumberMatchesBruteForce(t *testing.T) {
	for n := 0; n <= 300; n++ {
		got := digitOneInNumber(n)
		want := digitOneInNumberBrute(n)
		if got != want {
			t.Fatalf("digitOneInNumber(%d) = %d, brute force = %d", n, got, want)
		}
	}
}

func digitOneInNumberBrute(n int) int {
	cnt := 0
	for i := 1; i <= n; i++ {
		for x := i; x > 0; x /= 10 {
			if x%10 == 1 {
				cnt++
			}
		}
	}
	return cnt
}
