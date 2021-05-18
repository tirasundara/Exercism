package raindrops

import (
	"strconv"
)

// Convert number into a string that contains raindrop sounds
func Convert(input int) string {
	raindrops := ""

	if input%3 == 0 {
		raindrops += "Pling"
	}

	if input%5 == 0 {
		raindrops += "Plang"
	}

	if input%7 == 0 {
		raindrops += "Plong"
	}

	if raindrops == "" {
		return strconv.Itoa(input)
	}

	return raindrops
}
