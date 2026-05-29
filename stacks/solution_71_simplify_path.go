// Package stacks - LeetCode Problem 71: Simplify Path
// Given a string path, which is an absolute path (starting with a slash '/') to a file or directory,
// simplify it to the canonical path.
// Time: O(n), Space: O(n)
package stacks

// SimplifyPath returns a simplified canonical path.
func SimplifyPath(path string) string {
	stack := make([]string, 0)
	parts := split(path, '/')

	for _, part := range parts {
		switch part {
		case "", ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, part)
		}
	}

	result := ""
	for _, dir := range stack {
		result += "/" + dir
	}
	if result == "" {
		result = "/"
	}
	return result
}

func split(s string, delimiter byte) []string {
	var parts []string
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == delimiter {
			if start < i {
				parts = append(parts, s[start:i])
			}
			start = i + 1
		}
	}
	return parts
}