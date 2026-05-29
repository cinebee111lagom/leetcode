package linked_lists

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		vals  []int
		want  []int
	}{
		{"basic case", []int{1, 1, 2}, []int{1, 2}},
		{"all duplicates", []int{1, 1, 1}, []int{1}},
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"empty", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteDuplicates(buildList(tt.vals))
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("DeleteDuplicates() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}