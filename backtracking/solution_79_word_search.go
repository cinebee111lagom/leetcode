// Package backtracking - LeetCode Problem 79: Word Search
// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells (horizontally or vertically).
// Time: O(m*n * 4^L), Space: O(m*n)
package backtracking

// Exist determines if a word exists in the board.
func Exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board), len(board[0])

	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		if index == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[index] {
			return false
		}
		board[i][j] = '#'
		found := dfs(i+1, j, index+1) || dfs(i-1, j, index+1) ||
			dfs(i, j+1, index+1) || dfs(i, j-1, index+1)
		board[i][j] = word[index]
		return found
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}