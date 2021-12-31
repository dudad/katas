package sumofmultiples_test

import (
	. "github.com/dudad/katas/Go/sumofmultiples"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sprintf", func() {
	DescribeTable("FlagParser",
		func(in int, out int) {
			Expect(Multiple3And5(in)).To(Equal(out))
		},
		Entry("Positive Values Less Than 4", 3, 0),
		Entry("Value 4 - first with multiply of 3", 4, 3),
		Entry("Value 6 - first with multiply of 5", 6, 8),
		Entry("Value 15 - last before common multiply of 3 and 5", 15, 45),
		Entry("Value 16 - first with common multiply of 3 and 5", 16, 60),
		Entry("Negative value", -133, 0),
	)
})
