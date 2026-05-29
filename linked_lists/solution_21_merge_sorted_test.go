package linked_lists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		l1    []int
		l2    []int
		want  []int
	}{
		{"basic case", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"one empty", []int{}, []int{1}, []int{1}},
		{"both empty", []int{}, []int{}, []int{}},
		{"different lengths", []int{1, 3}, []int{2, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeTwoLists(buildList(tt.l1), buildList(tt.l2))
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}