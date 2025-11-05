package main

import (
	"testing"
)

func TestWordPuzzle(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		target   string
		expected bool
	}{
		{
			name: "基本功能 - 能找到单词（水平方向）",
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target:   "ABCCED",
			expected: true,
		},
		{
			name: "基本功能 - 能找到单词（垂直方向）",
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target:   "SEE",
			expected: true,
		},
		{
			name: "基本功能 - 找不到单词",
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target:   "ABCB",
			expected: false,
		},
		{
			name: "单字符网格 - 能找到",
			grid: [][]byte{
				{'A'},
			},
			target:   "A",
			expected: true,
		},
		{
			name: "单字符网格 - 找不到",
			grid: [][]byte{
				{'A'},
			},
			target:   "B",
			expected: false,
		},
		{
			name: "单行网格 - 能找到",
			grid: [][]byte{
				{'A', 'B', 'C'},
			},
			target:   "ABC",
			expected: true,
		},
		{
			name: "单列网格 - 能找到",
			grid: [][]byte{
				{'A'},
				{'B'},
				{'C'},
			},
			target:   "ABC",
			expected: true,
		},
		{
			name: "需要回溯的情况 - 能找到",
			grid: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			target:   "ABDC",
			expected: true,
		},
		{
			name: "需要回溯的情况 - 找不到（重复使用同一格）",
			grid: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			target:   "ABDCA",
			expected: false,
		},
		{
			name: "复杂路径 - 能找到",
			grid: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			target:   "ABEDGHI",
			expected: true,
		},
		{
			name: "复杂路径 - 找不到",
			grid: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			target:   "ABEDGHIA",
			expected: false,
		},
		{
			name: "多个起点 - 能找到",
			grid: [][]byte{
				{'A', 'A'},
				{'A', 'B'},
			},
			target:   "AAB",
			expected: true,
		},
		{
			name: "多个起点 - 找不到",
			grid: [][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			target:   "AAB",
			expected: false,
		},
		{
			name: "LCR 129 示例1 - 能找到",
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target:   "ABCCED",
			expected: true,
		},
		{
			name: "LCR 129 示例2 - 能找到",
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target:   "SEE",
			expected: true,
		},
		{
			name: "LCR 129 示例3 - 找不到",
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target:   "ABCB",
			expected: false,
		},
		{
			name: "长单词 - 能找到",
			grid: [][]byte{
				{'A', 'B', 'C', 'D'},
				{'E', 'F', 'G', 'H'},
				{'I', 'J', 'K', 'L'},
				{'M', 'N', 'O', 'P'},
			},
			target:   "ABCDHLPONMIEFG",
			expected: true,
		},
		{
			name: "长单词 - 找不到",
			grid: [][]byte{
				{'A', 'B', 'C', 'D'},
				{'E', 'F', 'G', 'H'},
				{'I', 'J', 'K', 'L'},
				{'M', 'N', 'O', 'P'},
			},
			target:   "ABCDHLPONMIEFGX",
			expected: false,
		},
		{
			name: "所有字符相同 - 能找到",
			grid: [][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			target:   "AAAA",
			expected: true,
		},
		{
			name: "所有字符相同 - 找不到（需要5个字符但只有4个格子）",
			grid: [][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			target:   "AAAAA",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建网格的副本，避免测试用例之间相互影响
			gridCopy := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]byte, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			result := wordPuzzle(gridCopy, tt.target)

			if result != tt.expected {
				t.Errorf("wordPuzzle() = %v, expected %v", result, tt.expected)
				t.Errorf("Grid: %v", tt.grid)
				t.Errorf("Target: %s", tt.target)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

