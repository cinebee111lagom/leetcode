// Package trees - LeetCode Problem 100: Same Tree
// Given the roots of two binary trees p and q, determine if they are the same.
// Time: O(n), Space: O(n)
package trees

// IsSameTree determines if two binary trees are the same.
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}