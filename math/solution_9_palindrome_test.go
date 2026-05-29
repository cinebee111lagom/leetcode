package math

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		x     int
		want  bool
	}{
		{"positive palindrome", 121, true},
		{"negative", -121, false},
		{"two digits", 10, false},
		{"single digit", 7, true},
		{"zero", 0, true},
		{"large palindrome", 1234321, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.x); got != tt.want {
				t.Errorf("IsPalindrome(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}