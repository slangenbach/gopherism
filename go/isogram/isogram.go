package isogram

import "strings"

// Check if a word is an isogram
func IsIsogram(word string) bool {
	lookup := make(map[rune]bool)

	for _, char := range strings.ToLower(word) {
		if lookup[char] {
			return false
		} else {
			if !(char == '-' || char == ' ') {
				lookup[char] = true
			}
		}
	}
	return true

}
