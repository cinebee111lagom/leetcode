// Package math - LeetCode Problem 9: Palindrome Number
// Given an integer x, return true if x is a palindrome, and false otherwise.
// Time: O(log n), Space: O(1)
package math

// IsPalindrome checks if an integer is a palindrome.
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}