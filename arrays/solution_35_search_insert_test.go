package arrays

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"basic case", []int{1, 3, 5, 6}, 5, 2},
		{"insert at beginning", []int{1, 3, 5, 6}, 0, 0},
		{"insert at end", []int{1, 3, 5, 6}, 7, 4},
		{"insert in middle", []int{1, 3, 5, 6}, 2, 1},
		{"single element found", []int{1}, 1, 0},
		{"single element insert", []int{1}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.nums, tt.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}