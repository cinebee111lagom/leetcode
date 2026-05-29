package trees

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name   string
		p      *TreeNode
		q      *TreeNode
		same   bool
	}{
		{
			"basic case",
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			true,
		},
		{
			"different structure",
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			&TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			false,
		},
		{
			"both empty",
			nil,
			nil,
			true,
		},
		{
			"one empty",
			&TreeNode{Val: 1},
			nil,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameTree(tt.p, tt.q); got != tt.same {
				t.Errorf("IsSameTree() = %v, want %v", got, tt.same)
			}
		})
	}
}