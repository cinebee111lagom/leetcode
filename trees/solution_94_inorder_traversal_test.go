package trees

import (
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name   string
		root   *TreeNode
		want   []int
	}{
		{
			"basic case",
			&TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			[]int{2, 1, 3},
		},
		{
			"single node",
			&TreeNode{Val: 1},
			[]int{1},
		},
		{
			"empty",
			nil,
			nil,
		},
		{
			"left skewed",
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{Val: 1},
				},
			},
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InorderTraversal(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
				return
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}