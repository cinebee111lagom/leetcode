package arrays

import "testing"

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"basic case", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"target at start", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{"target at end", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{"not found", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"single element", []int{1}, 1, 0},
		{"two elements", []int{3, 1}, 1, 1},
		{"no rotation", []int{1, 2, 3, 4, 5}, 3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchRotated(tt.nums, tt.target); got != tt.want {
				t.Errorf("SearchRotated() = %v, want %v", got, tt.want)
			}
		})
	}
}