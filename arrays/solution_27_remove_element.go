// Package arrays - LeetCode Problem 27: Remove Element
// Given an integer array nums and an integer val, remove all occurrences of val in-place.
// Return the number of elements in nums which are not equal to val.
// Time: O(n), Space: O(1)
package arrays

// RemoveElement removes all instances of val in-place and returns the new count.
func RemoveElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}