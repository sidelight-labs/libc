package killswitch

import (
	"sync"
	"time"
)

type KillSwitch struct {
	errStartedTime time.Time
	errThreshold   time.Duration
	mu             sync.Mutex
	errCount       int
	doneChan       chan bool
	C              chan bool
}

const maxErrorCountWithinThreshold = 10

func NewKillSwitch(errThreshold time.Duration) *KillSwitch {
	k := &KillSwitch{
		errStartedTime: time.Now(),
		errThreshold:   errThreshold,
		errCount:       0,
		doneChan:       make(chan bool),
		C:              make(chan bool, 1),
	}

	go k.start()
	return k
}

func (k *KillSwitch) Audit(_ error) {
	k.setErrCount(true)
}

func (k *KillSwitch) setErrCount(bump bool) {
	k.mu.Lock()
	defer k.mu.Unlock()

	if bump {
		k.errCount += 1
	} else {
		k.errCount = 0
		k.errStartedTime = time.Now()
	}
}

func (k *KillSwitch) getErrCount() int {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.errCount
}

func (k *KillSwitch) getErrStartedTime() time.Time {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.errStartedTime
}

func (k *KillSwitch) start() {
	ticker := time.NewTicker(k.errThreshold)

	defer close(k.C)

	for {
		select {
		case <-k.doneChan:
			ticker.Stop()
			return
		case <-ticker.C:
			k.setErrCount(false)
		default:
			if k.getErrCount() >= maxErrorCountWithinThreshold {
				if time.Now().Sub(k.getErrStartedTime()) <= k.errThreshold {
					k.C <- true
				}
			}
		}
	}
}

func (k *KillSwitch) Stop() {
	k.doneChan <- true
}
