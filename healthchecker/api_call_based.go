package healthchecker

import (
	"fmt"
	"net/http"
	"time"
)

type apiCallBasedHealthCheck struct {
	endpoint string
	*CommonHealthCheck
}

func NewApiCallBasedHealthCheck(endpoint string, interval time.Duration) HealthChecker {
	return &apiCallBasedHealthCheck{endpoint, &CommonHealthCheck{interval}}
}

func (hc *apiCallBasedHealthCheck) IsHealthy() bool {
	resp, err := http.Get(hc.endpoint)
	if err != nil {
		fmt.Printf("could not hit endpoint %s, err %v\n", hc.endpoint, err)
		return false
	}
	return resp.StatusCode == http.StatusOK
}
