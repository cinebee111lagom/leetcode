package math

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name  string
		x     int
		want  int
	}{
		{"4", 4, 2},
		{"8", 8, 2},
		{"0", 0, 0},
		{"1", 1, 1},
		{"16", 16, 4},
		{"100", 100, 10},
		{"large number", 2147395599, 46339},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySqrt(tt.x); got != tt.want {
				t.Errorf("MySqrt(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}