package lcr132

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
		{name: "需要取模 n=120", bambooLen: 120, expected: 953271190},
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

func TestCuttingBambooMatchesPowMod(t *testing.T) {
	for n := 2; n <= 80; n++ {
		got := cuttingBamboo(n)
		want := cuttingBambooByPowMod(n)
		if got != want {
			t.Fatalf("cuttingBamboo(%d) = %d, pow-mod = %d", n, got, want)
		}
	}
}

// 余数分类：n=3a+b 时尽量切成 3，余 1 改成 2+2，再快速幂取模。
func cuttingBambooByPowMod(n int) int {
	const mod int64 = 1000000007
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	a, b := n/3, n%3
	if b == 1 {
		return int(powMod(3, a-1, mod) * 4 % mod)
	}
	if b == 2 {
		return int(powMod(3, a, mod) * 2 % mod)
	}
	return int(powMod(3, a, mod))
}

func powMod(base, exp int, mod int64) int64 {
	ans := int64(1)
	x := int64(base) % mod
	for exp > 0 {
		if exp&1 == 1 {
			ans = ans * x % mod
		}
		x = x * x % mod
		exp >>= 1
	}
	return ans
}
