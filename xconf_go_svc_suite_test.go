package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestXconfGoSvc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "XconfGoSvc Suite")
}
