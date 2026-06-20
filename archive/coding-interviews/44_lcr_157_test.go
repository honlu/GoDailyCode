package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGoodsOrder(t *testing.T) {
	tests := []struct {
		name     string
		goods    string
		expected []string
	}{
		{
			name:     "空字符串",
			goods:    "",
			expected: []string{}, // 代码未处理空字符串情况，返回空切片
		},
		{
			name:     "单字符",
			goods:    "a",
			expected: []string{"a"},
		},
		{
			name:     "两个字符",
			goods:    "ab",
			expected: []string{"ab", "ba"},
		},
		{
			name:     "三个字符",
			goods:    "abc",
			expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name:     "有重复字符 - aab",
			goods:    "aab",
			expected: []string{"aab", "aba", "baa"},
		},
		{
			name:     "有重复字符 - aabb",
			goods:    "aabb",
			expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
		{
			name:     "四个不同字符",
			goods:    "abcd",
			expected: []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb", "bacd", "badc", "bcad", "bcda", "bdac", "bdca", "cabd", "cadb", "cbad", "cbda", "cdab", "cdba", "dabc", "dacb", "dbac", "dbca", "dcab", "dcba"},
		},
		{
			name:     "数字字符串",
			goods:    "123",
			expected: []string{"123", "132", "213", "231", "312", "321"},
		},
		{
			name:     "所有字符相同",
			goods:    "aaa",
			expected: []string{"aaa"},
		},
		{
			name:     "混合字符",
			goods:    "a1b",
			expected: []string{"a1b", "ab1", "1ab", "1ba", "ba1", "b1a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goodsOrder(tt.goods)
			
			// 对结果和期望值进行排序以便比较
			sort.Strings(result)
			sort.Strings(tt.expected)
			
			// 规范化空切片（nil 和 [] 在 reflect.DeepEqual 中不相等）
			if result == nil {
				result = []string{}
			}
			if tt.expected == nil {
				tt.expected = []string{}
			}
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("goodsOrder(%q) = %v, expected %v", tt.goods, result, tt.expected)
			} else {
				t.Logf("✓ Test passed: %s, 结果数量: %d", tt.name, len(result))
			}
		})
	}
}

// 测试边界情况
func TestGoodsOrderEdgeCases(t *testing.T) {
	t.Run("空字符串", func(t *testing.T) {
		result := goodsOrder("")
		// 代码未处理空字符串情况，返回空切片
		if len(result) != 0 {
			t.Errorf("空字符串返回空切片, got %v", result)
		}
	})

	t.Run("单字符", func(t *testing.T) {
		result := goodsOrder("x")
		if len(result) != 1 || result[0] != "x" {
			t.Errorf("单字符 'x' 应该返回 ['x'], got %v", result)
		}
	})

	t.Run("验证排列数量 - 无重复字符", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected int // 期望的排列数量 n!
		}{
			{"a", 1},      // 1! = 1
			{"ab", 2},     // 2! = 2
			{"abc", 6},    // 3! = 6
			{"abcd", 24},  // 4! = 24
		}

		for _, tc := range testCases {
			result := goodsOrder(tc.input)
			if len(result) != tc.expected {
				t.Errorf("goodsOrder(%q) 应该返回 %d 个排列, got %d", tc.input, tc.expected, len(result))
			}
		}
	})

	t.Run("验证去重功能", func(t *testing.T) {
		result := goodsOrder("aab")
		sort.Strings(result)
		expected := []string{"aab", "aba", "baa"}
		sort.Strings(expected)
		
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("goodsOrder('aab') 应该去重, got %v, expected %v", result, expected)
		}
		
		// 验证没有重复项
		seen := make(map[string]bool)
		for _, s := range result {
			if seen[s] {
				t.Errorf("结果中发现重复项: %s", s)
			}
			seen[s] = true
		}
	})
}

