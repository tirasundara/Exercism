package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is type map[string]int that stores word count
type Frequency map[string]int

// WordCount calculates word count in a sentence
func WordCount(s string) Frequency {
	f := make(Frequency)

	fieldFunc := func(r rune) bool {
		return unicode.IsSpace(r) || (r != '\'' && unicode.IsPunct(r)) || unicode.IsSymbol(r)
	}
	words := strings.FieldsFunc(s, fieldFunc)
	trimFunc := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}
	for _, word := range words {
		trimmed := strings.ToLower(strings.TrimFunc(word, trimFunc))
		if _, found := f[trimmed]; !found {
			f[trimmed] = 1
		} else {
			f[trimmed]++
		}
	}

	return f
}
