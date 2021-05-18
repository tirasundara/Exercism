package leap

// IsLeapYear reports wheter the year is a Leap year or not
func IsLeapYear(year int) bool {

	if year%4 != 0 {
		return false
	}

	if year%100 == 0 {
		return year%400 == 0
	}

	return true
}
