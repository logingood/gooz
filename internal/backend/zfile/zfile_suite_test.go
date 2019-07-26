package zfile_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestZfile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zfile Suite")
}
