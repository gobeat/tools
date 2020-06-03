package tools

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Common", func() {
	It("Must should not panic if no errors", func() {
		defer func() {
			Expect(recover()).To(BeNil())
		}()
		Must(nil)
	})

	It("Must should panic on errors", func() {
		defer func() {
			Expect(recover()).NotTo(BeNil())
		}()
		Must(nil, errors.New("there is an error"))
	})
})
