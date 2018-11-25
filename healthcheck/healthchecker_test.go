package healthcheck

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"
)

type MockSingleCheck struct {
	calls []bool
	call  int
}

func (mc *MockSingleCheck) SingleCheck() bool {
	if mc.call >= len(mc.calls) {
		panic(fmt.Sprintf("could not call mock for %d time, it has only %d values", mc.call+1, len(mc.calls)))
	}
	value := mc.calls[mc.call]
	fmt.Println("now yielding: ", value)
	mc.call++
	return value
}
func TestHealthChecks(t *testing.T) {
	testCases := []struct {
		passedThreshold  int
		failedThreshold  int
		singleCheckCalls []bool
		expectedHealth   []bool
		description      string
	}{
		{
			passedThreshold:  1,
			failedThreshold:  1,
			singleCheckCalls: []bool{true, false, true},
			expectedHealth:   []bool{true, false, true},
		},
	}
	for _, tc := range testCases {
		hc := &HealthCheck{
			PassedThreshold: tc.passedThreshold,
			FailedThreshold: tc.failedThreshold,
			Interval:        0, // does not matter, we mock it
			S:               &MockSingleCheck{calls: tc.singleCheckCalls},
			LogOutput:       bytes.NewBuffer(nil),
		}

		c := make(chan time.Time)
		// Mock `after` with `c` channel to make sure that
		// test is time-independent
		after = func(d time.Duration) <-chan time.Time {
			return c
		}
		ctx, cancel := context.WithCancel(context.Background())
		hc.Run(ctx)

		if hc.IsHealthy() != false {
			t.Fatal("expected to be unhealthy initially")
		}

		for i := 0; i < len(tc.singleCheckCalls); i++ {
			c <- time.Time{}
			time.Sleep(1 * time.Second)

			if hc.IsHealthy() != tc.expectedHealth[i] {
				t.Fatalf("run %d, expected healthy to be %t, but it wasn't", i, tc.expectedHealth[i])
			}
		}
		cancel()
	}
}
