package commondenominators_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/dudad/katas/tree/main/Go/commondenominators"
)

var _ = Describe("Commondenominators", func() {
	DescribeTable("Commondenominators all",
		func(in [][]int, out string) {
			Expect(ConvertFracts(in)).To(Equal(out))
		},
		Entry("Empty", [][]int{}, ""),
		Entry("single", [][]int{{1, 2}}, "(1,2)"),
		Entry("multi", [][]int{{1, 2}, {1, 3}, {1, 4}}, "(6,12)(4,12)(3,12)"),
	)
})

var _ = Describe("Commondenominators", func() {
	DescribeTable("Commondenominators - GCD",
		func(a int, b int, out int) {
			Expect(GCD(a, b)).To(Equal(out))
		},
		Entry("no GCD", 5, 6, 1),
		Entry("some GCD", 200, 300, 100),
		Entry("some GCD reorderd", 300, 200, 100),
	)
})

var _ = Describe("Commondenominators", func() {
	DescribeTable("Commondenominators - LCM",
		func(a int, b int, out int) {
			Expect(LCM(a, b)).To(Equal(out))
		},
		Entry("one of numder is LCM", 6, 12, 12),
		Entry("when no common GCD", 5, 6, 30),
		Entry("with some >1 GCD", 200, 300, 600),
		Entry("with some >1 GCD reoreder args", 300, 200, 600),
	)
})
