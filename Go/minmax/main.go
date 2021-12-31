package minmax

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	var min int = math.MaxInt32
	var max int = math.MinInt32
	words := strings.Fields(in)

	for _, word := range words {
		val, _ := strconv.Atoi(word)
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}
