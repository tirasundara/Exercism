package bob

import (
	"strings"
	"unicode"
)

// Hey returns string reponse based on passed remark
func Hey(remark string) string {

	if isForecefulQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isShouting(remark) {
		return "Whoa, chill out!"
	} else if isAsking(remark) {
		return "Sure."
	} else if isSilence(remark) {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func isShouting(s string) bool {

	if hasNoLetters(s) {
		return false
	}

	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}

	return true
}

func isAsking(s string) bool {
	return strings.HasSuffix(strings.Trim(s, " "), "?")
}

func isForecefulQuestion(s string) bool {
	return isShouting(s) && isAsking(s)
}

func hasNoLetters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isSilence(s string) bool {
	trimmed := strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	return len(trimmed) == 0
}
