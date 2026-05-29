// Package strings - LeetCode Problem 28: Implement strStr()
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
// Time: O(n+m), Space: O(m)
package strings

// StrStr implements the KMP algorithm to find needle in haystack.
func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	lps := computeLPS(needle)

	i, j := 0, 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return i - j
			}
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}

func computeLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	i := 1
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}