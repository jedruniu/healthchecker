package healthchecker

import "time"

type HealthChecker interface {
	IsHealthy() bool
	GetInterval() time.Duration
}

type CommonHealthCheck struct {
	interval time.Duration
}

func (hc *CommonHealthCheck) GetInterval() time.Duration {
	return hc.interval
}
