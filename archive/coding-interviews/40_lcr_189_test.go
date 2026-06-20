package main

import (
	"testing"
)

func TestMechanicalAccumulator(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		expected int
	}{
		{
			name:     "target为1",
			target:   1,
			expected: 1,
		},
		{
			name:     "target为5",
			target:   5,
			expected: 15, // 5+4+3+2+1 = 15
		},
		{
			name:     "target为10",
			target:   10,
			expected: 55, // 10+9+8+7+6+5+4+3+2+1 = 55
		},
		{
			name:     "target为0",
			target:   0,
			expected: 0,
		},
		{
			name:     "target为100",
			target:   100,
			expected: 5050, // 100*101/2 = 5050
		},
		{
			name:     "target为3",
			target:   3,
			expected: 6, // 3+2+1 = 6
		},
		{
			name:     "target为2",
			target:   2,
			expected: 3, // 2+1 = 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mechanicalAccumulator(tt.target)
			if result != tt.expected {
				t.Errorf("mechanicalAccumulator(%d) = %d, expected %d", tt.target, result, tt.expected)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

