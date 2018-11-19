package healthchecker

import "time"

type HealthChecker interface {
	IsHealthy() bool
	GetInterval() time.Duration
}

type CommonHealthCheck struct {
	name     string
	interval time.Duration
}

func (hc *CommonHealthCheck) GetInterval() time.Duration {
	return hc.interval
}

func (hc CommonHealthCheck) String() string {
	return hc.name
}
