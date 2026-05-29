package strings

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name   string
		strs   []string
		prefix string
	}{
		{"basic case", []string{"flower", "flow", "flight"}, "fl"},
		{"no common", []string{"dog", "racecar", "car"}, ""},
		{"single string", []string{"alone"}, "alone"},
		{"empty array", []string{}, ""},
		{"empty strings", []string{""}, ""},
		{"mixed lengths", []string{"interspecies", "interstellar", "interstate"}, "inters"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.strs); got != tt.prefix {
				t.Errorf("LongestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.prefix)
			}
		})
	}
}