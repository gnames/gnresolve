package gnresolve_test

import (
	"io/ioutil"
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGnresolve(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gnresolve Suite")
}

var _ = BeforeSuite(func() {
	log.SetOutput(ioutil.Discard)
})
