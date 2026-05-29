// Package binary_search - LeetCode Problem 4: Median of Two Sorted Arrays
// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// Time: O(log(m+n)), Space: O(1)
package binary_search

// FindMedianSortedArrays finds the median of two sorted arrays.
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m

	totalLen := m + n
	half := (totalLen + 1) / 2

	for left <= right {
		i := (left + right) / 2
		j := half - i

		if i < m && j > 0 && nums1[i] < nums2[j-1] {
			left = i + 1
		} else if i > 0 && j < n && nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			var maxLeft float64
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else if nums1[i-1] > nums2[j-1] {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = float64(nums2[j-1])
			}

			if (totalLen % 2) == 1 {
				return maxLeft
			}

			var minRight float64
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else if nums1[i] < nums2[j] {
				minRight = float64(nums1[i])
			} else {
				minRight = float64(nums2[j])
			}

			return (maxLeft + minRight) / 2.0
		}
	}
	return 0
}