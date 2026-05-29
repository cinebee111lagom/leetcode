// Package arrays - LeetCode Problem 75: Sort Colors
// Given an array nums with n objects colored red, white, or blue, sort them in-place
// so that objects of the same color are adjacent, with the colors in the order red, white, blue.
// Time: O(n), Space: O(1)
package arrays

// SortColors sorts the array with 0=red, 1=white, 2=blue using Dutch National Flag algorithm.
func SortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}