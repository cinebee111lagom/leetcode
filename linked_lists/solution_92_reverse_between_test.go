package linked_lists

import (
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		vals  []int
		left  int
		right int
		want  []int
	}{
		{"basic case", []int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{"reverse first two", []int{1, 2, 3, 4, 5}, 1, 2, []int{2, 1, 3, 4, 5}},
		{"reverse last two", []int{1, 2, 3, 4, 5}, 4, 5, []int{1, 2, 3, 5, 4}},
		{"single element", []int{1}, 1, 1, []int{1}},
		{"two elements", []int{1, 2}, 1, 2, []int{2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseBetween(buildList(tt.vals), tt.left, tt.right)
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("ReverseBetween() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}