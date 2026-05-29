// Package stacks - LeetCode Problem 20: Valid Parentheses
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid. An input string is valid if:
// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.
// Time: O(n), Space: O(n)
package stacks

// IsValid checks if a string of brackets is valid.
func IsValid(s string) bool {
	stack := make([]byte, 0)
	mapping := map[byte]byte{')': '(', '}': '{', ']': '['}

	for i := 0; i < len(s); i++ {
		if closing, ok := mapping[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}