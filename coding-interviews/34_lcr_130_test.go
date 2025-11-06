package main

import (
	"testing"
)

func TestWardrobeFinishing(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		cnt      int
		expected int
	}{
		{
			name:     "基本功能 - 小网格，cnt较大",
			m:        3,
			n:        3,
			cnt:      5,
			expected: 9,
		},
		{
			name:     "基本功能 - cnt较小，限制访问",
			m:        3,
			n:        3,
			cnt:      1,
			expected: 3,
		},
		{
			name:     "边界情况 - 单格网格，cnt=0",
			m:        1,
			n:        1,
			cnt:      0,
			expected: 1,
		},
		{
			name:     "边界情况 - 单格网格，cnt较小",
			m:        1,
			n:        1,
			cnt:      10,
			expected: 1,
		},
		{
			name:     "边界情况 - 单行网格",
			m:        1,
			n:        5,
			cnt:      3,
			expected: 4,
		},
		{
			name:     "边界情况 - 单列网格",
			m:        5,
			n:        1,
			cnt:      3,
			expected: 4,
		},
		{
			name:     "cnt=0 - 只能访问(0,0)",
			m:        5,
			n:        5,
			cnt:      0,
			expected: 1,
		},
		{
			name:     "cnt=1 - 可以访问digit和<=1的格子",
			m:        3,
			n:        3,
			cnt:      1,
			expected: 3,
		},
		{
			name:     "cnt=2 - 可以访问更多格子",
			m:        3,
			n:        3,
			cnt:      2,
			expected: 6,
		},
		{
			name:     "LeetCode示例 - 机器人的运动范围",
			m:        2,
			n:        3,
			cnt:      1,
			expected: 3,
		},
		{
			name:     "LeetCode示例 - 机器人的运动范围2",
			m:        3,
			n:        1,
			cnt:      0,
			expected: 1,
		},
		{
			name:     "较大网格 - cnt适中",
			m:        10,
			n:        10,
			cnt:      5,
			expected: 21,
		},
		{
			name:     "较大网格 - cnt较大，可访问大部分",
			m:        10,
			n:        10,
			cnt:      15,
			expected: 94,
		},
		{
			name:     "cnt限制严格 - 只能访问左上角",
			m:        5,
			n:        5,
			cnt:      0,
			expected: 1,
		},
		{
			name:     "cnt限制中等 - 可访问部分区域",
			m:        5,
			n:        5,
			cnt:      3,
			expected: 10,
		},
		{
			name:     "矩形网格 - 长宽不等",
			m:        4,
			n:        6,
			cnt:      3,
			expected: 10,
		},
		{
			name:     "矩形网格 - 长宽不等，cnt较大",
			m:        4,
			n:        6,
			cnt:      10,
			expected: 24,
		},
		{
			name:     "测试digit函数 - 两位数",
			m:        10,
			n:        10,
			cnt:      1,
			expected: 3,
		},
		{
			name:     "测试digit函数 - 三位数",
			m:        100,
			n:        100,
			cnt:      1,
			expected: 3,
		},
		{
			name:     "cnt刚好能访问到边界",
			m:        2,
			n:        2,
			cnt:      2,
			expected: 4,
		},
		{
			name:     "cnt刚好不能访问到边界",
			m:        2,
			n:        2,
			cnt:      1,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wardrobeFinishing(tt.m, tt.n, tt.cnt)

			if result != tt.expected {
				t.Errorf("wardrobeFinishing(%d, %d, %d) = %v, expected %v", tt.m, tt.n, tt.cnt, result, tt.expected)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

// 测试digit辅助函数
func TestDigit(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "单位数 - 0",
			input:    0,
			expected: 0,
		},
		{
			name:     "单位数 - 5",
			input:    5,
			expected: 5,
		},
		{
			name:     "两位数 - 10",
			input:    10,
			expected: 1,
		},
		{
			name:     "两位数 - 23",
			input:    23,
			expected: 5,
		},
		{
			name:     "两位数 - 99",
			input:    99,
			expected: 18,
		},
		{
			name:     "三位数 - 100",
			input:    100,
			expected: 1,
		},
		{
			name:     "三位数 - 123",
			input:    123,
			expected: 6,
		},
		{
			name:     "三位数 - 999",
			input:    999,
			expected: 27,
		},
		{
			name:     "四位数 - 1000",
			input:    1000,
			expected: 1,
		},
		{
			name:     "四位数 - 1234",
			input:    1234,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := digit(tt.input)

			if result != tt.expected {
				t.Errorf("digit(%d) = %v, expected %v", tt.input, result, tt.expected)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

