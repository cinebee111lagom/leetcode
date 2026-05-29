// Package arrays - LeetCode Problem 33: Search in Rotated Sorted Array
// There is an integer array nums sorted in ascending order (with distinct values).
// Prior to being passed to your function, nums is possibly rotated.
// Given the array nums after rotation and an integer target, return the index of target if it is in nums, or -1 otherwise.
// Time: O(log n), Space: O(1)
package arrays

// SearchRotated searches for target in a rotated sorted array.
func SearchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}