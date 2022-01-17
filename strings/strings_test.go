package strings_test

import (
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/sidelight-labs/libc/strings"
	"testing"
)

func TestUnitStrings(t *testing.T) {
	spec.Run(t, "Strings", testStrings, spec.Report(report.Terminal{}))
}

func testStrings(t *testing.T, when spec.G, it spec.S) {
	it.Before(func() {
		RegisterTestingT(t)
	})

	when("Shorten()", func() {
		inputString := "0123456789"

		it("returns the same string when its shorter than the maximum", func() {
			result := strings.Shorten(inputString, 10)
			Expect(result).To(HaveLen(10))
			Expect(result).To(Equal(inputString))
		})

		it("returns a shortened string when it exceeds the maxiumum", func() {
			result := strings.Shorten(inputString, 5)
			Expect(result).To(HaveLen(5))
			Expect(result).To(Equal("01234"))
		})
	})

	when("StringToArray()", func() {
		it("returns an empty array with an empty string as input", func() {
			result := strings.StringToArray("")
			Expect(result).Should(BeEmpty())
		})

		it("splits the string as expected", func() {
			result := strings.StringToArray("a,b,c,d")
			Expect(result).Should(HaveLen(4))
			Expect(result[0]).Should(Equal("a"))
		})
	})
}
