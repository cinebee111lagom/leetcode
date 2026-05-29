package stacks

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		maxArea  int
	}{
		{"basic case", []int{2, 1, 5, 6, 2, 3}, 10},
		{"all same", []int{2, 2, 2}, 6},
		{"increasing", []int{1, 2, 3, 4, 5}, 9},
		{"decreasing", []int{5, 4, 3, 2, 1}, 9},
		{"empty", []int{}, 0},
		{"single bar", []int{2}, 2},
		{"two bars", []int{1, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestRectangleArea(tt.heights); got != tt.maxArea {
				t.Errorf("LargestRectangleArea() = %v, want %v", got, tt.maxArea)
			}
		})
	}
}