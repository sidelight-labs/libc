package numbers_test

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/sidelight-labs/libc/numbers"
	"testing"
)

func TestUnitNumbers(t *testing.T) {
	spec.Run(t, "Numbers", testNumbers, spec.Report(report.Terminal{}))
}

func testNumbers(t *testing.T, when spec.G, it spec.S) {
	it.Before(func() {
		RegisterTestingT(t)
	})

	when("NumberToUint64()", func() {
		it("returns 0 for an empty string", func() {
			var input json.Number = ""
			expected := uint64(0)

			output, err := numbers.NumberToUint64(input)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(expected))
		})
		it("throws an error for an invalid number", func() {
			var input json.Number = "invalid"

			_, err := numbers.NumberToUint64(input)
			Expect(err).To(HaveOccurred())
		})
		it("parses a valid number as expected", func() {
			var input json.Number = "123"
			expected := uint64(123)

			output, err := numbers.NumberToUint64(input)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(expected))
		})
	})
}
