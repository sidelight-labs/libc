package logger_test

import (
	"errors"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/sidelight-labs/libc/logger"
	"testing"
)

func TestUnitLogger(t *testing.T) {
	spec.Run(t, "Logger", testLogger, spec.Report(report.Terminal{}))
}

func testLogger(t *testing.T, when spec.G, it spec.S) {
	it.Before(func() {
		RegisterTestingT(t)
	})

	when("Wrap()", func() {
		const (
			errorMsg = "the error that occurred"
			addition = "the additional message"
		)

		it("wraps error messages", func() {
			err := logger.Wrap(errors.New(errorMsg), addition)
			Expect(err.Error()).To(MatchRegexp("github.com/sidelight-labs/libc/logger_test.testLogger.func\\d+.\\d+\\[\\d+\\]: the additional message: the error that occurred"))
		})
	})
}
