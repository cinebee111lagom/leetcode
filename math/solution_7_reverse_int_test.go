package math

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		x     int
		want  int
	}{
		{"positive", 123, 321},
		{"negative", -123, -321},
		{"ends with zero", 120, 21},
		{"overflow", 2147483647, 0},
		{"underflow", -2147483648, 0},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.x); got != tt.want {
				t.Errorf("Reverse(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}