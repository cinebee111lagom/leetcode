// Package stacks - LeetCode Problem 84: Largest Rectangle in Histogram
// Given an array of integers heights representing the bar height of the histogram,
// find the largest rectangular area in the histogram.
// Time: O(n), Space: O(n)
package stacks

// LargestRectangleArea finds the largest rectangle area in the histogram.
func LargestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	stack := make([]int, 0)
	maxArea := 0

	for i := 0; i <= n; i++ {
		h := 0
		if i < n {
			h = heights[i]
		}
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}
	return maxArea
}