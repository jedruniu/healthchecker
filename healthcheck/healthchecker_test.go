package healthcheck

import (
	"bytes"
	"context"
	"testing"
	"time"
)

type MockSingleCheck struct {
	calls []bool
	call  int
}

func (mc *MockSingleCheck) SingleCheck() bool {
	value := mc.calls[mc.call]
	mc.call += 1
	return value
}
func TestHealthChecks(t *testing.T) {
	hc := &HealthCheck{
		PassedThreshold: 1,
		FailedThreshold: 1,
		Interval:        1 * time.Nanosecond, //does not matter
		S:               &MockSingleCheck{calls: []bool{true}},
		LogOutput:       bytes.NewBufferString(""),
	}
	c := make(chan time.Time)
	after = func(d time.Duration) <-chan time.Time {
		return c
	}
	ctx, cancel := context.WithCancel(context.Background())
	hc.Run(ctx)

	if hc.IsHealthy() != false {
		t.Fatal("expected to be unhealthy initially")
	}
	c <- time.Time{}
	cancel()

	if hc.IsHealthy() != true {
		t.Fatal("expected to be healthy after one passed healthcheck")
	}
}
