package errors_test

import (
	. "github.com/hiroaki-yamamoto/yakuza/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ValueError", func() {
	var err *ValueError
	BeforeEach(func() {
		err = ValueError{}.Init("Test", "This is a test")
	})
	It("Should explain what happened.", func() {
		Expect(err.Error()).Should(Equal(err.Message))
	})
})
