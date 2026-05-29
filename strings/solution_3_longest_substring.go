// Package strings - LeetCode Problem 3: Longest Substring Without Repeating
// Given a string s, find the length of the longest substring without repeating characters.
// Time: O(n), Space: O(1)
package strings

// LengthOfLongestSubstring finds the length of the longest substring without repeating chars.
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	charIndex := make(map[byte]int)
	maxLen := 0
	start := 0
	for i := 0; i < len(s); i++ {
		if idx, ok := charIndex[s[i]]; ok && idx >= start {
			start = idx + 1
		}
		charIndex[s[i]] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}