package anagram

import "sort"

func sortedRunesFromString(str string) string {
	r := []rune{}
	for _, char := range str {
		r = append(r, char)
	}
	sort.Slice(r, func(i int, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func Anagrams(word string, words []string) []string {
	if len(word) == 0 || len(words) == 0 {
		return nil
	}
	sortedWord := sortedRunesFromString(word)
	ret := []string{}

	for _, candidate := range words {
		if sortedWord == sortedRunesFromString(candidate) {
			ret = append(ret, candidate)
		}
	}
	if len(ret) == 0 {
		return nil
	}

	return ret
}
