package arrays

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name    string
		height  []int
		wantMax int
	}{
		{"basic case", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"two elements", []int{1, 1}, 1},
		{"decreasing", []int{4, 3, 2, 1, 4}, 16},
		{"single line", []int{1, 2}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea(tt.height); got != tt.wantMax {
				t.Errorf("MaxArea() = %v, want %v", got, tt.wantMax)
			}
		})
	}
}