package healthchecker

import "time"

type Healthchecker interface {
	IsHealthy() bool
	GetInterval() time.Duration
}
