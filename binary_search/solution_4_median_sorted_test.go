package binary_search

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		nums2  []int
		median float64
	}{
		{"basic case", []int{1, 3}, []int{2}, 2.0},
		{"even length", []int{1, 2}, []int{3, 4}, 2.5},
		{"different sizes", []int{1, 3, 5, 7}, []int{2, 4, 6, 8}, 4.5},
		{"empty first", []int{}, []int{1}, 1.0},
		{"single each", []int{1}, []int{2}, 1.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMedianSortedArrays(tt.nums1, tt.nums2); got != tt.median {
				t.Errorf("FindMedianSortedArrays() = %v, want %v", got, tt.median)
			}
		})
	}
}