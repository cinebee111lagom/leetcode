package dp

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		maxSum  int
	}{
		{"basic case", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"single element", []int{1}, 1},
		{"all positive", []int{1, 2, 3}, 6},
		{"all negative", []int{-1, -2, -3}, -1},
		{"mixed", []int{-2, -1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.nums); got != tt.maxSum {
				t.Errorf("MaxSubArray() = %d, want %d", got, tt.maxSum)
			}
		})
	}
}