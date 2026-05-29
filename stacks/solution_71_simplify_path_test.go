package stacks

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		want  string
	}{
		{"basic case", "/home/", "/home"},
		{"up level", "/home/foo/", "/home/foo"},
		{"double dot", "/../", "/"},
		{"current dir", "/home//foo/", "/home/foo"},
		{"root", "/", "/"},
		{"multiple up", "/a/b/../c/", "/a/c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimplifyPath(tt.path); got != tt.want {
				t.Errorf("SimplifyPath(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}