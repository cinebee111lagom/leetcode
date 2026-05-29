package backtracking

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{"n=1", 1, []string{"()"}},
		{"n=2", 2, []string{"(())", "()()"}},
		{"n=3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateParenthesis(tt.n)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("GenerateParenthesis(%d) = %v, want %v", tt.n, got, tt.expected)
			}
		})
	}
}