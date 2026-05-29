package linked_lists

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name  string
		vals  []int
		want  []int
	}{
		{"basic case", []int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{"odd length", []int{1, 2, 3}, []int{2, 1, 3}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"single element", []int{1}, []int{1}},
		{"empty", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SwapPairs(buildList(tt.vals))
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("SwapPairs() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}