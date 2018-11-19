package healthchecker

import (
	"fmt"
	"net/http"
	"time"
)

type apiCallBasedHealthChecker struct {
	endpoint string
	interval time.Duration
}

func NewApiCallBasedHealthChecker(endpoint string, interval time.Duration) HealthChecker {
	return &apiCallBasedHealthChecker{endpoint, interval}
}

func (hc *apiCallBasedHealthChecker) IsHealthy() bool {
	resp, err := http.Get(hc.endpoint)
	if err != nil {
		fmt.Printf("could not hit endpoint %s, err %v\n", hc.endpoint, err)
		return false
	}
	return resp.StatusCode == http.StatusOK
}

func (hc *apiCallBasedHealthChecker) GetInterval() time.Duration {
	return hc.interval
}
