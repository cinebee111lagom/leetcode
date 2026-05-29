package arrays

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		want  []int
		count int
	}{
		{"basic case", []int{1, 1, 2}, []int{1, 2}, 2},
		{"all duplicates", []int{1, 1, 1}, []int{1}, 1},
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}, 3},
		{"empty", []int{}, []int{}, 0},
		{"single element", []int{1}, []int{1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicates(tt.nums)
			if got != tt.count {
				t.Errorf("RemoveDuplicates() count = %v, want %v", got, tt.count)
			}
			if !reflect.DeepEqual(tt.nums[:got], tt.want) {
				t.Errorf("RemoveDuplicates() array = %v, want %v", tt.nums[:got], tt.want)
			}
		})
	}
}