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

	)
})
