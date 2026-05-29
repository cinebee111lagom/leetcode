// Package math - LeetCode Problem 69: Sqrt(x)
// Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
// Time: O(log n), Space: O(1)
package math

// MySqrt returns the integer square root of x.
func MySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}