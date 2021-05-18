package grains

import (
	"errors"
)

// Square returns 2^(n-1). And also error will be returned if n is out of range (1..64)
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("out of range")
	}

	if n == 1 {
		return 1, nil
	}

	return uint64(2 << (n - 2)), nil
}

// Total returns sum of the Square(n). Where n from 1 to 64
func Total() uint64 {
	sum := uint64(0)

	for n := 1; n <= 64; n++ {
		sq, err := Square(n)
		if err != nil {
			return 0
		}
		sum += sq
	}
	return sum
}
