package tools

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{3, 3, 6},
		{4, 4, 8},
	}

	for _, tt := range tests {
		if got := Sum(tt.a, tt.b); got != tt.want {
			t.Errorf("Sum(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestTranspose(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]float64
		expected [][]float64
	}{
		{
			name:     "2x2 matrix",
			input:    [][]float64{{1, 2}, {3, 4}},
			expected: [][]float64{{1, 3}, {2, 4}},
		},
		{
			name:     "3x3 matrix",
			input:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]float64{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Transpose(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Transpose(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
