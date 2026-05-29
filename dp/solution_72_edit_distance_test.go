package dp

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name   string
		word1  string
		word2  string
		dist   int
	}{
		{"horse to ros", "horse", "ros", 3},
		{"intention to execution", "intention", "execution", 5},
		{"empty to abc", "", "abc", 3},
		{"abc to empty", "abc", "", 3},
		{"same word", "abc", "abc", 0},
		{"single char diff", "abcd", "abce", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDistance(tt.word1, tt.word2); got != tt.dist {
				t.Errorf("MinDistance(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.dist)
			}
		})
	}
}