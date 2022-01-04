package factorialdecomposition

import (
	"bytes"
	"fmt"
)

func Decomp(n int) string {
	primesDecomp := map[int]int{}
	primes := []int{}
	for factor := 2; factor <= n; factor++ {
		num := factor
		for divisor := 2; num > 1; {
			if num % divisor == 0 {
				primesDecomp[divisor] += 1
				if( primesDecomp[divisor] == 1) {
					primes = append(primes, divisor)
				}
				num /= divisor
			} else {
				divisor++
			}
		}
	}

	var buffer bytes.Buffer
	for _, prime := range primes {
		exp := primesDecomp[prime]
		buffer.WriteString(fmt.Sprintf(" * %d", prime))
		if exp > 1 {
			buffer.WriteString(fmt.Sprintf("^%d", exp))
		}
	} 

	return buffer.String()[3:]
}
