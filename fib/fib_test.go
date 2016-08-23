package fib_test

import (
	. "public/fib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fib", func() {
	It("Gives the nth fibonacci number", func() {
		Expect(Fib(4)).To(Equal(2))
	})
})
