package arrays

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		val   int
		count int
	}{
		{"basic case", []int{3, 2, 2, 3}, 3, 2},
		{"no occurrences", []int{1, 2, 3}, 4, 3},
		{"all occurrences", []int{3, 3, 3}, 3, 0},
		{"empty", []int{}, 1, 0},
		{"single element match", []int{1}, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.nums, tt.val); got != tt.count {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.count)
			}
		})
	}
}