package controller_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var tReference *testing.T

func TestController(t *testing.T) {
	tReference = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}
