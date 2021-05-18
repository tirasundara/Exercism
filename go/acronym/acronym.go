package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	splitter := func(c rune) bool {
		return unicode.IsSpace(c) || c == '-'
	}
	words := strings.FieldsFunc(s, splitter)
	abbr := ""
	f := func(c rune) bool {
		return unicode.IsLetter(c)
	}

	for _, word := range words {
		i := strings.IndexFunc(word, f)
		if i > -1 {
			abbr += strings.ToUpper(string(word[i]))
		}
	}

	return abbr
}
