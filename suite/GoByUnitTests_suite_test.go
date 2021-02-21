package gobyunittests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gobyunittests/fourth_sample"
	// "gobyunittests/main"
	"gobyunittests/second_sample"
	"gobyunittests/sixth_example"
	"gobyunittests/third_sample"
)

func TestGoByUnitTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoByUnitTests Suite")
}
