package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"basic case", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"different order", []int{3, 2, 4}, 6, []int{1, 2}},
		{"same value", []int{3, 3}, 6, []int{0, 1}},
		{"larger array", []int{1, 5, 3, 7, 11}, 10, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}