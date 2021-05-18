package romannumerals

// ToRomanNumeral converts
func ToRomanNumeral(n int) (string, error) {
	num := [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	sym := [13]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	i := len(num) - 1
	roman := ""

	for {

		if n == 0 {
			break
		}

		div := n / num[i]
		n = n % num[i]

		for {

			if div == 0 {
				break
			}

			roman += sym[i]
			div--
		}
		i--
	}

	return roman, nil
}
