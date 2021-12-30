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
	for _, char := range roman {
		val := decodeDigit(char)
		sum += val
	}
	if roman == "I" {
		return 1
	}
	return sum
}

// func main() {
// 	Decode("II")
// }
