// Package arrays - LeetCode Problem 80: Remove Duplicates from Sorted Array II
// Given an integer array nums sorted in non-decreasing order, remove duplicates in-place
// such that each unique element appears at most twice. Return the new length.
// Time: O(n), Space: O(1)
package arrays

// RemoveDuplicatesII removes duplicates allowing at most twice.
func RemoveDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	count := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[count-2] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}