package xconf_go_svc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestXconfGoSvc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "XconfGoSvc Suite")
}
