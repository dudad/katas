package roman

// Very popular Kata
// Converting Roman Literals to decimal

func decodeDigit(char rune) int {
	switch char {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	default:
		return 1000
	}
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
	return sum
}
