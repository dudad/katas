package bookseller

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	sum := 0
	for _, book := range listArt {
		valStr := strings.Fields(book)[1]
		val, _ := strconv.Atoi(valStr)
		sum += val
	}
	return fmt.Sprintf("(A - %d)", sum)
}
