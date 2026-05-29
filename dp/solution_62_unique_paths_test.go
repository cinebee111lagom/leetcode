package dp

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{"3x7 grid", 3, 7, 28},
		{"3x2 grid", 3, 2, 3},
		{"1x1 grid", 1, 1, 1},
		{"2x2 grid", 2, 2, 2},
		{"1x10 grid", 1, 10, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("UniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}