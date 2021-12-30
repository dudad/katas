package bookseller

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func categoryMapToString(categoryMap map[string]int, listCat []string) string {
	if len(categoryMap) == 0 || len(listCat) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, cat := range listCat {
		buffer.WriteString(fmt.Sprintf(" - (%s : %d)", cat, categoryMap[cat]))
	}
	return buffer.String()[3:]
}

func pareBookString(book string) (string, int) {
	items := strings.Fields(book)
	name := items[0]
	valStr := items[1]
	val, _ := strconv.Atoi(valStr)

	return name[0:1], val
}

func buildCategoryMap(listArt []string) map[string]int {
	categoryMap := make(map[string]int)
	for _, book := range listArt {
		category, val := pareBookString(book)
		categoryMap[category] += val
	}
	return categoryMap
}

func StockList(listArt []string, listCat []string) string {
	return categoryMapToString(buildCategoryMap(listArt), listCat)
}
