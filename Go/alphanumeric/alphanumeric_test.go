package alphanumeric_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/dudad/katas/commits/main/Go/alphanumeric"
)

var _ = Describe("Alphanumeric", func() {
	DescribeTable("Alphanumeric all",
		func(word string, out bool) {
			Expect(Alphanumeric(word)).To(Equal(out))
		},
		Entry("Empty", "", false),
		Entry("other chars", ".*?", false),
		Entry("single letter", "a", true),
	)
})
