package strings

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		name      string
		haystack  string
		needle    string
		wantIndex int
	}{
		{"basic case", "hello", "ll", 2},
		{"not found", "aaaaa", "bba", -1},
		{"empty needle", "hello", "", 0},
		{"needle at start", "hello", "he", 0},
		{"needle at end", "hello", "lo", 3},
		{"empty haystack", "", "a", -1},
		{"same strings", "a", "a", 0},
		{"no match", "abc", "d", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.haystack, tt.needle); got != tt.wantIndex {
				t.Errorf("StrStr(%q, %q) = %d, want %d", tt.haystack, tt.needle, got, tt.wantIndex)
			}
		})
	}
}