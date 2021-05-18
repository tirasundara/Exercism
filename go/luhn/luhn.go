// Package luhn contains luhn-realted algorithm and validator
package luhn

import (
	"strconv"
	"strings"
)

// Valid returns true wheter is valid string based on Luhn algorithm
func Valid(s string) bool {
	if len(strings.TrimSpace(s)) < 2 {
		return false
	}

	s = strings.Join(strings.Fields(s), "")
	runes := reverseRunes([]rune(s))
	sum := 0

	for i, r := range runes {
		p, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}

		if i%2 != 0 {
			p *= 2
			if p > 9 {
				p -= 9
			}
		}
		sum += p
	}

	return sum%10 == 0
}

func reverseRunes(r []rune) []rune {
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r
}
