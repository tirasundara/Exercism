package strain

// Ints is an array of integers
type Ints []int

// Strings is an array of strings
type Strings []string

// Lists is an array of array integers
type Lists [][]int

// Keep Ints
func (ints Ints) Keep(f func(int) bool) Ints {
	var res Ints

	for _, n := range ints {
		if f(n) {
			res = append(res, n)
		}
	}

	return res
}

// Discard Ints
func (ints Ints) Discard(f func(int) bool) Ints {
	var res Ints

	for _, n := range ints {
		if !f(n) {
			res = append(res, n)
		}
	}

	return res
}

// Keep Strings
func (strings Strings) Keep(f func(string) bool) Strings {
	var res Strings

	for _, n := range strings {
		if f(n) {
			res = append(res, n)
		}
	}

	return res
}

// Discard Strings
func (strings Strings) Discard(f func(string) bool) Strings {
	var res Strings

	for _, n := range strings {
		if !f(n) {
			res = append(res, n)
		}
	}

	return res
}

// Keep Lists
func (lists Lists) Keep(f func([]int) bool) Lists {
	var res Lists

	for _, n := range lists {
		if f(n) {
			res = append(res, n)
		}
	}

	return res
}

// Discard Lists
func (lists Lists) Discard(f func([]int) bool) Lists {
	var res Lists

	for _, n := range lists {
		if !f(n) {
			res = append(res, n)
		}
	}

	return res
}
