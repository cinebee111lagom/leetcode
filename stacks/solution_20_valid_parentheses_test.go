package stacks

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		valid bool
	}{
		{"basic case", "()", true},
		{"nested", "()[]{}", true},
		{"mixed", "([])", true},
		{"invalid order", "(]", false},
		{"invalid close", "([)]", false},
		{"empty string", "", true},
		{"unclosed open", "((", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.s); got != tt.valid {
				t.Errorf("IsValid(%q) = %v, want %v", tt.s, got, tt.valid)
			}
		})
	}
}