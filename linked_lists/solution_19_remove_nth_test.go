package linked_lists

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name  string
		vals  []int
		n     int
		want  []int
	}{
		{"remove second", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"remove first", []int{1, 2}, 1, []int{1}},
		{"remove last", []int{1, 2}, 2, []int{2}},
		{"single element", []int{1}, 1, []int{}},
		{"two elements remove first", []int{1, 2}, 2, []int{2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.vals)
			got := RemoveNthFromEnd(head, tt.n)
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}