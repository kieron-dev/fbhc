package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Search", func() {

	DescribeTable("contains", func(needle, haystack string, expected bool) {
		Expect(Contains(needle, haystack)).To(Equal(expected))
	},
		Entry("xyzxzyx", "xyzxzyx", "xyzxyzxzyx", false),
		Entry("ababac", "ababac", "abababac", false),
		Entry("abac", "abac", "ababac", false),
		Entry("kier", "kier", "fookieronfoo", true),
		Entry("abcabd", "abcabd", "abcabcabd", false),
	)

	DescribeTable("counter ex", func(needle, counter string) {
		Expect(MakeCounterExample(needle)).To(Equal(counter))
		Expect(Contains(needle, MakeCounterExample(needle))).To(BeFalse())
	},
		Entry("xyzxzyx", "xyzxzyx", "xyzxyzxzyx"),
		Entry("x", "x", "Impossible"),
		Entry("abac", "abac", "ababac"),
	)

})
