package healthchecker

import (
	"fmt"
	"net/http"
)

type apiCallBasedHealthCheck struct {
	endpoint string
}

func NewApiCallBased(endpoint string) SingleChecker {
	return &apiCallBasedHealthCheck{endpoint}
}

func (hc *apiCallBasedHealthCheck) SingleCheck() bool {
	resp, err := http.Get(hc.endpoint)
	if err != nil {
		fmt.Printf("could not hit endpoint %s, err %v\n", hc.endpoint, err)
		return false
	}
	return resp.StatusCode == http.StatusOK
}
