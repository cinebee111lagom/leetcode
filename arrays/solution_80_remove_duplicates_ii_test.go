package arrays

import "testing"

func TestRemoveDuplicatesII(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		want  int
	}{
		{"basic case", []int{1, 1, 1, 2, 2, 3}, 5},
		{"already within limit", []int{1, 1, 2, 2}, 4},
		{"all same", []int{1, 1, 1, 1}, 2},
		{"no duplicates", []int{1, 2, 3}, 3},
		{"two elements", []int{1, 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicatesII(tt.nums); got != tt.want {
				t.Errorf("RemoveDuplicatesII() = %v, want %v", got, tt.want)
			}
		})
	}
}