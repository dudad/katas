package anagram_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/dudad/katas/tree/main/anagram"
)

var _ = Describe("Anagram", func() {
	DescribeTable("Anagram all",
		func(word string, words []string, out []string) {
			Expect(Anagrams(word, words)).To(Equal(out))
		},
		Entry("Empty", "wow", []string{}, nil),
		Entry("All are anagrams", "wow", []string{"oww", "wwo"}, []string{"oww", "wwo"}),
	)
})
