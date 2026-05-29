// Package backtracking - LeetCode Problem 22: Generate Parentheses
// Given n pairs of parentheses, generate all combinations of well-formed parentheses.
// Time: O(4^n / sqrt(n)), Space: O(n)
package backtracking

// GenerateParenthesis generates all valid parenthesis combinations.
func GenerateParenthesis(n int) []string {
	var result []string
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, current string, open int, close int, n int) {
	if len(current) == n*2 {
		*result = append(*result, current)
		return
	}
	if open < n {
		backtrack(result, current+"(", open+1, close, n)
	}
	if close < open {
		backtrack(result, current+")", open, close+1, n)
	}
}