package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var tReference *testing.T

func TestService(t *testing.T) {
	tReference = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}
