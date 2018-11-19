package healthchecker

import "time"

type HealthChecker interface {
	IsHealthy() bool
	GetInterval() time.Duration
}
