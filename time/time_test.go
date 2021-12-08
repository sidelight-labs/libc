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
}
