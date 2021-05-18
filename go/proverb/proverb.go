package proverb

const (
	start  = "For want of a "
	middle = " the "
	end    = " was lost."
	last   = "And all for the want of a "
	stop   = "."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	length := len(rhyme)

	if length == 0 {
		return []string{}
	}

	var proverbs []string
	var line string
	for i := 0; i < length; i++ {
		if i+1 < length {
			line = start + rhyme[i] + middle + rhyme[i+1] + end
		} else {
			line = last + rhyme[0] + stop
		}

		proverbs = append(proverbs, line)
	}

	return proverbs
}
