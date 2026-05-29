package trees

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name   string
		root   *TreeNode
		valid  bool
	}{
		{
			"basic case",
			&TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			true,
		},
		{
			"invalid",
			&TreeNode{
				Val: 5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			false,
		},
		{
			"single node",
			&TreeNode{Val: 1},
			true,
		},
		{
			"empty",
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBST(tt.root); got != tt.valid {
				t.Errorf("IsValidBST() = %v, want %v", got, tt.valid)
			}
		})
	}
}