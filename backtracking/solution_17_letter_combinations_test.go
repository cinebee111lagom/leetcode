package backtracking

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{"23", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"empty", "", []string{}},
		{"2", "2", []string{"a", "b", "c"}},
		{"7", "7", []string{"p", "q", "r", "s"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LetterCombinations(tt.digits)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("LetterCombinations(%q) = %v, want %v", tt.digits, got, tt.expected)
			}
		})
	}
}