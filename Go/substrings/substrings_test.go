package substrings_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/dudad/katas/commits/main/Go/substrings"
)

var _ = Describe("Substrings", func() {
	DescribeTable("Alphanumeric all",
		func(arr1 []string, arr2 []string, out []string) {
			Expect(InArray(arr1, arr2)).To(Equal(out))
		},
		Entry("basic", []string{"live", "arp", "strong"},
			[]string{"lively", "alive", "harp", "sharp", "armstrong"},
			[]string{"arp", "live", "strong"}),
		Entry("none", []string{"tarp", "mice", "bull"},
			[]string{"lively", "alive", "harp", "sharp", "armstrong"},
			[]string{}),
	)
})
