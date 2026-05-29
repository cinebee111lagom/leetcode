// Package trees - LeetCode Problem 94: Binary Tree Inorder Traversal
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
// Time: O(n), Space: O(n)
package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InorderTraversal returns the inorder traversal of a binary tree.
func InorderTraversal(root *TreeNode) []int {
	var result []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result
}