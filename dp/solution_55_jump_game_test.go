package dp

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"basic case", []int{2, 3, 1, 1, 4}, true},
		{"cannot reach", []int{3, 2, 1, 0, 4}, false},
		{"single element", []int{0}, true},
		{"two elements", []int{2, 0}, true},
		{"large jump", []int{1, 0, 2}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanJump(tt.nums); got != tt.want {
				t.Errorf("CanJump() = %v, want %v", got, tt.want)
			}
		})
	}
}