package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is type map[string]int that stores word count
type Frequency map[string]int

// WordCount calculates word count in a sentence
func WordCount(phrase string) Frequency {
	f := make(Frequency)

	re := regexp.MustCompile(`\w+('\w)*`)
	for _, word := range re.FindAllString(strings.ToLower(phrase), -1) {
		f[word]++
	}
	return f
}
