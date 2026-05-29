// Package math - LeetCode Problem 7: Reverse Integer
// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range, return 0.
// Time: O(log(x)), Space: O(1)
package math

// Reverse reverses an integer.
func Reverse(x int) int {
	result := 0
	for x != 0 {
		digit := x % 10
		if result > 2147483647/10 || (result == 2147483647/10 && digit > 7) {
			return 0
		}
		if result < -2147483648/10 || (result == -2147483648/10 && digit < -8) {
			return 0
		}
		result = result*10 + digit
		x /= 10
	}
	return result
}