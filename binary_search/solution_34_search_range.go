// Package binary_search - LeetCode Problem 34: Find First and Last Position
// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
// Time: O(log n), Space: O(1)
package binary_search

// SearchRange finds the start and end position of target in nums.
func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := findLeftBound(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := findRightBound(nums, target)
	return []int{left, right}
}

func findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func findRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2 + 1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}