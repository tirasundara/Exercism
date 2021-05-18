package pangram

import (
	"unicode"
)

const alphSize = 26

// IsPangram returns true or false wheter the passed s is pangram sentence or not
func IsPangram(s string) bool {
	if len(s) < alphSize {
		return false
	}

	alphabetPresence := make(map[rune]bool, alphSize)
	for _, r := range s {
		if unicode.IsLetter(r) {
			alphabetPresence[unicode.ToLower(r)] = true
		}
	}

	return len(alphabetPresence) == alphSize
}
