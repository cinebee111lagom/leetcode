// Package arrays - LeetCode Problem 26: Remove Duplicates from Sorted Array
// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place
// such that each unique element appears only once. Return the number of unique elements.
// Time: O(n), Space: O(1)
package arrays

// RemoveDuplicates removes duplicates in-place and returns the count of unique elements.
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	unique := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[unique] = nums[i]
			unique++
		}
	}
	return unique
}