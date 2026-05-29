// Package stacks - LeetCode Problem 32: Longest Valid Parentheses
// Given a string containing just the characters '(' and ')', find the length of the longest valid parentheses substring.
// Time: O(n), Space: O(1)
package stacks

// LongestValidParentheses finds the longest valid parentheses substring.
func LongestValidParentheses(s string) int {
	maxLen := 0
	left, right := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLen = max(maxLen, left*2)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLen = max(maxLen, left*2)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}