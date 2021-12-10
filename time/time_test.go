package time_test

import (
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/sidelight-labs/libc/time"
	"testing"
)

func TestUnitTime(t *testing.T) {
	spec.Run(t, "Time", testTime, spec.Report(report.Terminal{}))
}

func testTime(t *testing.T, when spec.G, it spec.S) {
	it.Before(func() {
		RegisterTestingT(t)
	})

	when("TimeToEpoch()", func() {
		it("throws an error when the time is not compatible", func() {
			inputString := "abc"
			_, err := time.TimeToEpoch(inputString)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(fmt.Sprintf(time.TimeParsingError, inputString)))
		})

		it("parses a YYYY-MM-DD formatted input string as expected", func() {
			inputString := "2021-02-19"
			result, err := time.TimeToEpoch(inputString)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(1613692800))
		})

		it("parses a RFC-3339 formatted input string as expected", func() {
			inputString := "2021-02-19T00:00:00Z"
			result, err := time.TimeToEpoch(inputString)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(1613692800))
		})
	})

	when("FormatTime", func() {
		it("throws an error when the format is incorrect", func() {
			incorrect := "noooo"
			_, err := time.FormatTime(incorrect)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(fmt.Sprintf(time.InvalidTime, incorrect)))
		})
		it("throws parses the input as expected", func() {
			input := "2021-05-22"
			output, err := time.FormatTime(input)

			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal("May 22 2021"))
		})
	})

	when("TimestampToDate", func() {
		it("parses the input as expected", func() {
			input := "2021-11-10T17:50:45Z"

			result, err := time.TimestampToDate(input)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal("2021-11-10"))
		})
		it("throws an error when the input is malformed", func() {
			input := "malformed"

			_, err := time.FormatTime(input)
			Expect(err).To(HaveOccurred())
		})
	})
}
