package lcr185

import (
	"math"
	"testing"
)

const probabilityTolerance = 1e-6

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < probabilityTolerance
}

func TestStatisticsProbability(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected []float64
	}{
		{
			name:     "1个骰子",
			num:      1,
			expected: []float64{1.0 / 6, 1.0 / 6, 1.0 / 6, 1.0 / 6, 1.0 / 6, 1.0 / 6},
		},
		{
			name: "官方示例 2个骰子",
			num:  2,
			expected: []float64{
				1.0 / 36, 2.0 / 36, 3.0 / 36, 4.0 / 36, 5.0 / 36, 6.0 / 36,
				5.0 / 36, 4.0 / 36, 3.0 / 36, 2.0 / 36, 1.0 / 36,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := statisticsProbability(tt.num)
			if len(got) != len(tt.expected) {
				t.Fatalf("statisticsProbability(%d) len = %d, expected %d", tt.num, len(got), len(tt.expected))
			}
			for i := range got {
				if !floatEqual(got[i], tt.expected[i]) {
					t.Errorf("statisticsProbability(%d)[%d] = %v, expected %v", tt.num, i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestStatisticsProbabilityMatchesBruteForce(t *testing.T) {
	for n := 1; n <= 4; n++ {
		got := statisticsProbability(n)
		want := statisticsProbabilityBrute(n)
		if len(got) != len(want) {
			t.Fatalf("n=%d len = %d, brute force len = %d", n, len(got), len(want))
		}
		sum := 0.0
		for i := range got {
			if !floatEqual(got[i], want[i]) {
				t.Fatalf("n=%d sum=%d: got %v, brute force %v", n, n+i, got[i], want[i])
			}
			sum += got[i]
		}
		if !floatEqual(sum, 1.0) {
			t.Fatalf("n=%d probabilities sum to %v, expected 1", n, sum)
		}
	}
}

func statisticsProbabilityBrute(n int) []float64 {
	total := 1
	for i := 0; i < n; i++ {
		total *= 6
	}
	ways := make([]int, 5*n+1)
	var dfs func(left, sum int)
	dfs = func(left, sum int) {
		if left == 0 {
			ways[sum-n]++
			return
		}
		for face := 1; face <= 6; face++ {
			dfs(left-1, sum+face)
		}
	}
	dfs(n, 0)
	out := make([]float64, len(ways))
	for i, w := range ways {
		out[i] = float64(w) / float64(total)
	}
	return out
}
