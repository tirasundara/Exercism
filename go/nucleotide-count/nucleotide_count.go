package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts is a method on the DNA type. A method is a function with a special receiver argument.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, r := range d {
		if _, ok := h[r]; !ok {
			return h, errors.New("invalid strand found")
		}
		h[r]++
	}
	return h, nil
}
