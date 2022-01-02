package pyramid_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPyramid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pyramid Suite")
}
