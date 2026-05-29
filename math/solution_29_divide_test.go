package math

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		want  int
	}{
		{"positive", 10, 3, 3},
		{"negative dividend", -10, 3, -3},
		{"negative divisor", 10, -3, -3},
		{"both negative", -10, -3, 3},
		{"exact division", 100, 10, 10},
		{"overflow edge", -2147483648, -1, 2147483647},
		{"zero dividend", 0, 1, 0},
		{"one", 5, 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.a, tt.b); got != tt.want {
				t.Errorf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}