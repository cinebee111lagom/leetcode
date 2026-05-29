// Package math - LeetCode Problem 29: Divide Two Integers
// Given two integers dividend and divisor, divide two integers without using multiplication, division, or mod operator.
// Return the quotient after dividing the dividend by the divisor.
// Time: O(log n), Space: O(1)
package math

// Divide divides two integers without using multiplication/division.
func Divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	negative := (dividend < 0) != (divisor < 0)
	a, b := abs(dividend), abs(divisor)
	quotient := 0

	for a >= b {
		shift := 0
		for a >= b<<(shift+1) {
			shift++
		}
		quotient += 1 << shift
		a -= b << shift
	}

	if negative {
		return -quotient
	}
	return quotient
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}