package triangle

import "math"

// Kind defines the type of triangle
type Kind int

const (
	// NaT means not a triangle
	NaT Kind = iota // 0
	// Equ means equilateral triangle
	Equ // 1
	// Iso means isosceles triangle
	Iso // 2
	// Sca means scalene triangle
	Sca // 3
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	var sum float64 = 0
	sides := [3]float64{a, b, c}

	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
		sum += side
	}

	for _, side := range sides {
		if side > (sum - side) {
			return NaT
		}
	}

	if a == b && a == c && b == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else if a != b && a != c && b != c {
		k = Sca
	}

	return k
}
