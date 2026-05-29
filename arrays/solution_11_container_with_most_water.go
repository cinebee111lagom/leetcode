// Package arrays - LeetCode Problem 11: Container With Most Water
// Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
// n vertical lines are drawn such that the two endpoints of line i are at (i, ai) and (i, 0).
// Find two lines that together with the x-axis form a container that holds the most water.
package arrays

// MaxArea returns the maximum area that can be contained.
// Time: O(n), Space: O(1)
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		width := right - left
		h := min(height[left], height[right])
		maxArea = max(maxArea, width*h)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}