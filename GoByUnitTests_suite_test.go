package GoByUnitTests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoByUnitTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoByUnitTests Suite")
}
