package pyramid_test

import (
	. "github.com/dudad/katas/tree/main/Go/pyramid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pyramid", func() {
	DescribeTable("Pyramid all",
		func(in int, out [][]int) {
			Expect(Pyramid(in)).To(Equal(out))
		},
		Entry("Empty", 0, [][]int{}),
		Entry("signle", 1, [][]int{{1}, }),
		Entry("multiply", 3, [][]int{{1}, {1,1}, {1,1,1} }),
	)
})
