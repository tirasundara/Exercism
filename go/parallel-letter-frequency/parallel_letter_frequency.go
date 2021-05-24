package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates the letter frequency concurrently
func ConcurrentFrequency(list []string) FreqMap {
	c := make(chan FreqMap)

	for _, s := range list {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

  result := FreqMap{}
	for range list {
		for letter, freq := range <-c {
			result[letter] += freq
		}
	}

	return result
}
