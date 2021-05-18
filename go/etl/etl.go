package etl

import (
	"strings"
)

// Transform transforms legacy scrable score into a new one
func Transform(scrableScore map[int][]string) map[string]int {
	var transformed = make(map[string]int)

	for k, v := range scrableScore {
		for _, letter := range v {
			transformed[strings.ToLower(letter)] = k
		}
	}

	return transformed
}
