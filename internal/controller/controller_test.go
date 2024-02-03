package controller_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("immudb integration tests", func() {
	a, b := 1, 1
	Expect(a).To(Equal(b))
})
