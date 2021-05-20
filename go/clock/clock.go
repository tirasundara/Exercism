package clock

import (
	"fmt"
	"math"
)

const (
	hours = 24
	minutes = 60
	hLower = 0
	hUpper = 23
	mLower = 0
	mUpper = 59
)

// Clock respresents the clock that hold hour, and minute
type Clock struct {
	h, m int
}

// New returns Clock type
func New(h, m int) Clock {
	if m < mLower {
		mToh := math.Ceil(math.Abs(float64(m)) / float64(minutes))
		h = h - int(mToh)
		m = pmod(m, minutes)
	}

	div, mod := divmod(m, minutes)
	_, h = divmod(h + div, hours)
	m = mod

	return Clock{h, m}
}

// String prints the clock in string mode
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add adds m-minute to the clock
func (c Clock) Add(m int) Clock {
	return New(c.h, c.m + m)
}

// Subtract substracts m-minute to the clock
func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m - m)
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = pmod(numerator, denominator)
	return
}

// Positive modulo, returns non negative solution to x % d  
func pmod(x, d int) int {  
  x = x % d
  if x >= 0 { return x }
  if d < 0 { return x - d }
  return x + d
}

