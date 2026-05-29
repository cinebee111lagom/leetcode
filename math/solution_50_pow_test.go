package math

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	tests := []struct {
		name  string
		x     float64
		n     int
		want  float64
	}{
		{"2^10", 2.0, 10, 1024.0},
		{"2.1^3", 2.1, 3, 9.261},
		{"x^0", 5.0, 0, 1.0},
		{"negative power", 2.0, -2, 0.25},
		{"0^positive", 0.0, 5, 0.0},
		{"1^any", 1.0, 100, 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.x, tt.n); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Pow(%v, %d) = %v, want %v", tt.x, tt.n, got, tt.want)
			}
		})
	}
}