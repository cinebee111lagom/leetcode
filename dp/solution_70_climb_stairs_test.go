package dp

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"2 steps", 2, 2},
		{"3 steps", 3, 3},
		{"1 step", 1, 1},
		{"4 steps", 4, 5},
		{"5 steps", 5, 8},
		{"0 steps", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.n); got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}