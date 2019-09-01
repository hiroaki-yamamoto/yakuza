package series_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSeries(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Series Suite")
}
