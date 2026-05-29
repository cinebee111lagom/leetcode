// Package dp - LeetCode Problem 55: Jump Game
// Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Determine if you can reach the last index.
// Time: O(n), Space: O(1)
package dp

// CanJump determines if you can reach the last index.
func CanJump(nums []int) bool {
	maxReach := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
	}
	return true
}