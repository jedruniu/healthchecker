// healthchecker package provides a funtionality to activelly monitor any target (e.g. file, website).
package healthchecker

import (
	"context"
	"fmt"
	"time"
)

type SingleChecker interface {
	SingleCheck() bool
}

type HealthCheck struct {
	// Unique check name.
	Name string
	// Thresholds define the level of healthy/unhealthy.
	FailedThreshold int
	PassedThreshold int
	// Data polling interval.
	Interval time.Duration
	// A target to monitor.
	S SingleChecker

	failedCount int
	passedCount int
	healthy     bool
}

func (hc *HealthCheck) IsHealthy() bool { return hc.healthy }

func (hc *HealthCheck) Run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-time.After(hc.Interval):
				hc.RunSingleCheck()
				fmt.Printf("Name: %q\t Health: %v\n", hc.Name, hc.IsHealthy())
			case <-ctx.Done():
				fmt.Printf("Name: %q\t Context terminated\n", hc.Name)
				return
			}
		}
	}()
}

func (hc *HealthCheck) RunSingleCheck() {
	if sc := hc.S.SingleCheck(); sc {
		hc.passedCount++
		hc.failedCount = 0
	} else {
		hc.failedCount++
		hc.passedCount = 0
	}
	if hc.passedCount >= hc.PassedThreshold {
		hc.healthy = true
	}
	if hc.failedCount >= hc.FailedThreshold {
		hc.healthy = false
	}
}
