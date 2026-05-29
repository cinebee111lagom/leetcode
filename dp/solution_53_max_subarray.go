// Package dp - LeetCode Problem 53: Maximum Subarray
// Given an integer array nums, find the subarray with the largest sum, and return its sum.
// Time: O(n), Space: O(1)
package dp

// MaxSubArray finds the maximum sum of a contiguous subarray.
func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(currentSum+nums[i], nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}