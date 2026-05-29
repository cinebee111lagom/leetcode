// Package backtracking - LeetCode Problem 17: Letter Combinations of Phone Number
// Given a string containing digits from 2-9, return all possible letter combinations that the number could represent.
// Time: O(4^n), Space: O(n)
package backtracking

var phone = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

// LetterCombinations returns all letter combinations from phone number input.
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var result []string
	backtrackLetters(&result, "", digits, 0)
	return result
}

func backtrackLetters(result *[]string, current string, digits string, index int) {
	if len(current) == len(digits) {
		*result = append(*result, current)
		return
	}
	letters := phone[digits[index]]
	for _, letter := range letters {
		backtrackLetters(result, current+letter, digits, index+1)
	}
}