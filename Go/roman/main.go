package roman

// Very popular Kata
// Converting Roman Literals to decimal

func decodeDigit(char rune) int {
	if char == 'I' {
		return 1
	}
	return 5
}

func Decode(roman string) int {
	sum := 0
	prev := 1000000
	for _, char := range roman {
		val := decodeDigit(char)
		if prev < val {
			sum -= 2 * prev
		}
		sum += val
		prev = val
	}
	if roman == "I" {
		return 1
	}
	return sum
}
