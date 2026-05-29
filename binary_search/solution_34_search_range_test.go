package binary_search

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"basic case", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"not found", []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{"single element found", []int{1}, 1, []int{0, 0}},
		{"single element not found", []int{1}, 0, []int{-1, -1}},
		{"all same", []int{1, 1, 1, 1}, 1, []int{0, 3}},
		{"two elements", []int{1, 2}, 1, []int{0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchRange(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}