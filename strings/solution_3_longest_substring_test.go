package strings

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		want  int
	}{
		{"abcabcbb", "abcabcbb", 3},
		{"bbbbb", "bbbbb", 1},
		{"pwwkew", "pwwkew", 3},
		{"empty", "", 0},
		{"single", "a", 1},
		{"dvdf", "dvdf", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}