// Package strings - LeetCode Problem 14: Longest Common Prefix
// Write a function to find the longest common prefix string among an array of strings.
// Time: O(n*k), Space: O(1)
package strings

// LongestCommonPrefix finds the longest common prefix among an array of strings.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && len(strs[i]) < len(prefix) || prefix != strs[i][:len(prefix)] {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			break
		}
	}
	return prefix
}