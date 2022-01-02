package commondenominators

import (
	"bytes"
	"fmt"
)

// GCD - Greatest common divisor
func GCD(a int, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

// LCM - Least common multiple
func LCM(a int, b int) int {
	return (a / GCD(a, b)) * b
}

func ConvertFracts(a [][]int) string {
	if len(a) == 0 {
		return ""
	}

	lcm := 1
	for _, pair := range a {
		lcm = LCM(lcm, pair[1])
	}
	var buffer bytes.Buffer

	for _, pair := range a {

		N := pair[0] * (lcm / pair[1])
		D := pair[1] * (lcm / pair[1])
		buffer.WriteString(fmt.Sprintf("(%d,%d)", N, D))
	}

	return buffer.String()
}
