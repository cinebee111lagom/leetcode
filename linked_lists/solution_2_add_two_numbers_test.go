package linked_lists

import (
	"reflect"
	"testing"
)

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	dummy := &ListNode{}
	current := dummy
	for _, v := range vals {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return dummy.Next
}

func listToSlice(l *ListNode) []int {
	if l == nil {
		return []int{}
	}
	var result []int
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name  string
		l1    []int
		l2    []int
		want  []int
	}{
		{"basic case", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"carry over", []int{9, 9, 9, 9}, []int{9, 9, 9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 1}},
		{"different lengths", []int{1}, []int{9, 9, 9}, []int{0, 0, 0, 1}},
		{"zeros", []int{0}, []int{0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := buildList(tt.l1)
			l2 := buildList(tt.l2)
			got := AddTwoNumbers(l1, l2)
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}