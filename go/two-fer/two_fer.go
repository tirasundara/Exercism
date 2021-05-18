package twofer

import "fmt"

// ShareWith receives one string argument called name
// it will returns string with the following format:
// "One for {name}, one for me."
// but, if name is empty, then the following string will be returned:
// "One for you, one for me."
func ShareWith(name string) string {

	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
