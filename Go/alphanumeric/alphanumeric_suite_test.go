package alphanumeric_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAlphanumeric(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Alphanumeric Suite")
}
