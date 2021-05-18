package accumulate

// Accumulate collection to a new collection by applying passed func
func Accumulate(coll []string, f func(s string) string) []string {
	accumulated := []string{}

	for _, v := range coll {
		accumulated = append(accumulated, f(v))
	}

	return accumulated
}
