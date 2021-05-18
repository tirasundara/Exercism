package isogram

import (
	"unicode"
)

// IsIsogram returns true or false wheter the string s is an isogram or not
func IsIsogram(s string) bool {

	// initialize charCount map
	charCount := map[rune]int{}

	for _, r := range s {

		// skip non letter
		if !unicode.IsLetter(r) {
			continue
		}

		char := unicode.ToLower(r)
		count := charCount[char]

		if count > 0 {
			return false
		}

		charCount[char]++
	}

	return true
}
