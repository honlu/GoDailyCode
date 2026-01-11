package main

import (
	"math"
	"testing"
)

const tolerance = 1e-9 // 浮点数比较容差

// floatEqual 比较两个浮点数是否相等（在容差范围内）
func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < tolerance
}

func TestMyPow(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{
		{
			name:     "n为0",
			x:        2.0,
			n:        0,
			expected: 1.0,
		},
		{
			name:     "n为1",
			x:        2.0,
			n:        1,
			expected: 2.0,
		},
		{
			name:     "n为2",
			x:        2.0,
			n:        2,
			expected: 4.0,
		},
		{
			name:     "n为3",
			x:        2.0,
			n:        3,
			expected: 8.0,
		},
		{
			name:     "n为4",
			x:        2.0,
			n:        4,
			expected: 16.0,
		},
		{
			name:     "n为10",
			x:        2.0,
			n:        10,
			expected: 1024.0,
		},
		{
			name:     "n为负数 -1",
			x:        2.0,
			n:        -1,
			expected: 0.5,
		},
		{
			name:     "n为负数 -2",
			x:        2.0,
			n:        -2,
			expected: 0.25,
		},
		{
			name:     "n为负数 -3",
			x:        2.0,
			n:        -3,
			expected: 0.125,
		},
		{
			name:     "x为0",
			x:        0.0,
			n:        5,
			expected: 0.0,
		},
		{
			name:     "x为1",
			x:        1.0,
			n:        100,
			expected: 1.0,
		},
		{
			name:     "x为-1, n为偶数",
			x:        -1.0,
			n:        2,
			expected: 1.0,
		},
		{
			name:     "x为-1, n为奇数",
			x:        -1.0,
			n:        3,
			expected: -1.0,
		},
		{
			name:     "x为小数 0.5, n为2",
			x:        0.5,
			n:        2,
			expected: 0.25,
		},
		{
			name:     "x为小数 2.0, n为-2",
			x:        2.0,
			n:        -2,
			expected: 0.25,
		},
		{
			name:     "x为小数 2.1, n为3",
			x:        2.1,
			n:        3,
			expected: 9.261,
		},
		{
			name:     "x为小数 2.0, n为-10",
			x:        2.0,
			n:        -10,
			expected: 0.0009765625,
		},
		{
			name:     "x为小数 0.00001, n为2147483647",
			x:        0.00001,
			n:        2147483647,
			expected: 0.0, // 非常小的数的大次幂接近0
		},
		{
			name:     "x为2.0, n为-2147483648",
			x:        2.0,
			n:        -2147483648,
			expected: 0.0, // 非常大的负数次幂接近0
		},
		{
			name:     "x为1.0, n为-2147483648",
			x:        1.0,
			n:        -2147483648,
			expected: 1.0,
		},
		{
			name:     "x为-2.0, n为2",
			x:        -2.0,
			n:        2,
			expected: 4.0,
		},
		{
			name:     "x为-2.0, n为3",
			x:        -2.0,
			n:        3,
			expected: -8.0,
		},
		{
			name:     "x为-2.0, n为-2",
			x:        -2.0,
			n:        -2,
			expected: 0.25,
		},
		{
			name:     "x为-2.0, n为-3",
			x:        -2.0,
			n:        -3,
			expected: -0.125,
		},
		{
			name:     "x为3.0, n为5",
			x:        3.0,
			n:        5,
			expected: 243.0,
		},
		{
			name:     "x为3.0, n为-5",
			x:        3.0,
			n:        -5,
			expected: 1.0 / 243.0,
		},
		{
			name:     "x为10.0, n为5",
			x:        10.0,
			n:        5,
			expected: 100000.0,
		},
		{
			name:     "x为10.0, n为-5",
			x:        10.0,
			n:        -5,
			expected: 0.00001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := myPow(tt.x, tt.n)
			if !floatEqual(result, tt.expected) {
				t.Errorf("myPow(%f, %d) = %f, expected %f", tt.x, tt.n, result, tt.expected)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

// 测试边界情况
func TestMyPowEdgeCases(t *testing.T) {
	t.Run("n为0，x为任意值", func(t *testing.T) {
		testCases := []float64{0.0, 1.0, 2.0, -1.0, 0.5, 100.0}
		for _, x := range testCases {
			result := myPow(x, 0)
			if !floatEqual(result, 1.0) {
				t.Errorf("myPow(%f, 0) = %f, expected 1.0", x, result)
			}
		}
	})

	t.Run("x为0，n为正数", func(t *testing.T) {
		testCases := []int{1, 2, 3, 10, 100}
		for _, n := range testCases {
			result := myPow(0.0, n)
			if !floatEqual(result, 0.0) {
				t.Errorf("myPow(0.0, %d) = %f, expected 0.0", n, result)
			}
		}
	})

	t.Run("x为1，n为任意值", func(t *testing.T) {
		testCases := []int{0, 1, -1, 2, -2, 100, -100, 2147483647, -2147483648}
		for _, n := range testCases {
			result := myPow(1.0, n)
			if !floatEqual(result, 1.0) {
				t.Errorf("myPow(1.0, %d) = %f, expected 1.0", n, result)
			}
		}
	})

	t.Run("x为-1，n为偶数", func(t *testing.T) {
		testCases := []int{2, 4, 6, 10, 100, -2, -4, -10}
		for _, n := range testCases {
			result := myPow(-1.0, n)
			if !floatEqual(result, 1.0) {
				t.Errorf("myPow(-1.0, %d) = %f, expected 1.0", n, result)
			}
		}
	})

	t.Run("x为-1，n为奇数", func(t *testing.T) {
		testCases := []int{1, 3, 5, 7, 101, -1, -3, -5, -101}
		for _, n := range testCases {
			result := myPow(-1.0, n)
			if !floatEqual(result, -1.0) {
				t.Errorf("myPow(-1.0, %d) = %f, expected -1.0", n, result)
			}
		}
	})

	t.Run("验证奇偶性处理", func(t *testing.T) {
		// 测试奇数n
		x := 2.0
		resultOdd := myPow(x, 5)
		expectedOdd := 32.0
		if !floatEqual(resultOdd, expectedOdd) {
			t.Errorf("奇数n处理失败: myPow(%f, 5) = %f, expected %f", x, resultOdd, expectedOdd)
		}

		// 测试偶数n
		resultEven := myPow(x, 6)
		expectedEven := 64.0
		if !floatEqual(resultEven, expectedEven) {
			t.Errorf("偶数n处理失败: myPow(%f, 6) = %f, expected %f", x, resultEven, expectedEven)
		}
	})

	t.Run("大n值测试", func(t *testing.T) {
		// 测试大n值，验证不会超时或内存溢出
		result := myPow(2.0, 30)
		expected := math.Pow(2.0, 30)
		if !floatEqual(result, expected) {
			t.Errorf("大n值测试失败: myPow(2.0, 30) = %f, expected %f", result, expected)
		}
	})
}
