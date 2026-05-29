// Package trees - LeetCode Problem 98: Validate Binary Search Tree
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
// Time: O(n), Space: O(n)
package trees

// IsValidBST determines if a binary tree is a valid BST.
func IsValidBST(root *TreeNode) bool {
	var validate func(node, min, max *TreeNode) bool
	validate = func(node, min, max *TreeNode) bool {
		if node == nil {
			return true
		}
		if min != nil && node.Val <= min.Val {
			return false
		}
		if max != nil && node.Val >= max.Val {
			return false
		}
		return validate(node.Left, min, node) && validate(node.Right, node, max)
	}
	return validate(root, nil, nil)
}