// Package arrays - LeetCode Problem 15: 3Sum
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that
// i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// The solution set must not contain duplicate triplets.
package arrays

import "sort"

// ThreeSum returns all unique triplets that sum to zero.
// Time: O(n^2), Space: O(1) excluding output
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		target := -nums[i]
		for left < right {
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}