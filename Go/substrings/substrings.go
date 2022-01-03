package substrings

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	ret := map[string]bool{}
	for _, word := range array1 {
		for _, el := range array2 {
			if strings.Contains(el, word) {
				ret[word] = true
				break
			}
		}
	}
	retArr := []string{}
	for key,_ := range ret {
		retArr = append(retArr,key)
	}
	sort.Strings(retArr)
	return retArr
}
