package clock

import (
	"fmt"
)

const (
	hoursPerday = 24
	minutesPerHour = 60
)

// Clock respresents the clock that hold hour, and minute
type Clock struct {
	h, m int
}

// New returns Clock type
func New(h, m int) Clock {
	return normalize(Clock{h, m})
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

func normalize(c Clock) Clock {
  c.h += c.m / minutesPerHour
  if c.m %= minutesPerHour; c.m < 0 {
  	c.h--
  	c.m += minutesPerHour
  }

  if c.h %= hoursPerday; c.h < 0 {
  	c.h += hoursPerday
  }
	return c
}

