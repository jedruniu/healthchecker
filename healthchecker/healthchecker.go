package healthchecker

import (
	"time"
)

type HealthChecker interface {
	IsHealthy() bool
	GetInterval() time.Duration
	RunSingleCheck()
}

type SingleChecker interface {
	SingleCheck() bool
}

type HealthCheck struct {
	Name            string
	FailedCnt       int
	FailedThreshold int
	PassedCnt       int
	PassedThreshold int
	Healthy         bool
	Interval        time.Duration
	S               SingleChecker
}

func NewHealthCheck(name string, failedThreshold, passedThreshold int, interval time.Duration, sc SingleChecker) HealthChecker {
	return &HealthCheck{name, 0, failedThreshold, 0, passedThreshold, false, interval, sc}
}

func (hc HealthCheck) GetInterval() time.Duration {
	return hc.Interval
}

func (hc HealthCheck) String() string {
	return hc.Name
}

func (hc HealthCheck) IsHealthy() bool { return hc.Healthy }

func (hc *HealthCheck) RunSingleCheck() {
	sc := hc.S.SingleCheck()
	if sc {
		hc.PassedCnt++
		hc.FailedCnt = 0
	} else {
		hc.FailedCnt++
		hc.PassedCnt = 0
	}
	if hc.PassedCnt >= hc.PassedThreshold {
		hc.Healthy = true
	}
	if hc.FailedCnt >= hc.FailedThreshold {
		hc.Healthy = false
	}
}
