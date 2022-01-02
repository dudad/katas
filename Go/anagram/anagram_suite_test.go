package anagram_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAnagram(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Anagram Suite")
}
