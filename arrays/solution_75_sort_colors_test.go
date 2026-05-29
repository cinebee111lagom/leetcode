package arrays

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		want  []int
	}{
		{"basic case", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"all same", []int{1, 1, 1}, []int{1, 1, 1}},
		{"reverse order", []int{2, 1, 0}, []int{0, 1, 2}},
		{"two colors", []int{2, 0, 1}, []int{0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortColors(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("SortColors() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}