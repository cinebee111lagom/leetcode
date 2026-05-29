package stacks

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		want  int
	}{
		{"basic case", "(()", 2},
		{"valid string", ")()())", 4},
		{"empty", "", 0},
		{"all open", "(((", 0},
		{"all close", ")))", 0},
		{"mixed", "()(()()", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestValidParentheses(tt.s); got != tt.want {
				t.Errorf("LongestValidParentheses(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}