package factorialdecomposition_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/dudad/katas/tree/main/Go/factorialdecomposition"
)

var _ = Describe("Factorialdecomposition", func() {
	DescribeTable("Factorialdecomposition all",
		func(n int, out string) {
			Expect(Decomp(n)).To(Equal(out))
		},
		Entry("smallest", 2, "2"),
		Entry("some", 17, "2^15 * 3^6 * 5^3 * 7^2 * 11 * 13 * 17"),
		Entry("some", 5, "2^3 * 3 * 5"),
		Entry("some", 22, "2^19 * 3^9 * 5^4 * 7^3 * 11^2 * 13 * 17 * 19"),
		Entry("some", 14, "2^11 * 3^5 * 5^2 * 7^2 * 11 * 13"),
		Entry("some", 25, "2^22 * 3^10 * 5^6 * 7^3 * 11^2 * 13 * 17 * 19 * 23"),
	)
})
