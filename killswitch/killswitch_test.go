package killswitch_test

import (
	"errors"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/sidelight-labs/libc/killswitch"
)

func TestUnitKillSwitch(t *testing.T) {
	spec.Run(t, "KillSwitch Test", testKillSwitch, spec.Report(report.Terminal{}))
}

func testKillSwitch(t *testing.T, when spec.G, it spec.S) {
	it.Before(func() {
		RegisterTestingT(t)
	})

	it("performs the lifecycle", func() {
		k := killswitch.NewKillSwitch(500 * time.Millisecond)
		defer k.Stop()

		// throw errors, but they're spaced out
		// validate that the killswitch doesn't go off
		go func() {
			time.Sleep(100 * time.Millisecond)
			for range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11} {
				time.Sleep(time.Millisecond)
				k.Audit(errors.New("some-error"))
			}
		}()

		var result bool
		Consistently(k.C).ShouldNot(Receive(&result))

		// throw errors, but the threshold is broken
		// validate that the killswitch goes off
		go func() {
			for range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11} {
				k.Audit(errors.New("some-error"))
			}
		}()

		Eventually(k.C).Should(Receive(&result))
	})
}
