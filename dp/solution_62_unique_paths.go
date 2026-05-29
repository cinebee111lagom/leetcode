// Package dp - LeetCode Problem 62: Unique Paths
// A robot is located at the top-left corner of a m x n grid.
// The robot can only move either down or right at any point in time.
// How many different paths are there to reach the bottom-right corner?
// Time: O(m*n), Space: O(m*n)
package dp

// UniquePaths returns the number of unique paths in an m x n grid.
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}