package healthcheck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	FailedThreshold int    `json:"failedThreshold"`
	PassedThreshold int    `json:"passedThreshold"`
	Interval        int    `json:"interval"`
	Target          string `json:"target"`
}

func HealthChecksFromConfig(cfg *[]Config) []HealthChecker {
	var healthChecks []HealthChecker

	for _, c := range *cfg {
		// create common part of Healthcheck
		check := &HealthCheck{
			Name:            c.Name,
			FailedThreshold: c.FailedThreshold,
			PassedThreshold: c.PassedThreshold,
			Interval:        time.Duration(c.Interval) * time.Second,
			LogOutput:       os.Stdout,
		}

		// dispatch SingleChecker dynamically by key in config
		var s func(string) SingleChecker

		switch c.Type {
		case "redis_based":
			s = NewRedisBased
		case "file_based":
			s = NewFileBased
		case "shell_based":
			s = NewShellBased
		case "api_call_based":
			s = NewApiCallBased
		default:
			panic(fmt.Sprintf("unknown checker detected: %s", c.Type))
		}
		check.S = s(c.Target)

		healthChecks = append(healthChecks, check)
	}
	return healthChecks
}

func ReadConfig(path string) (*[]Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c []Config
	err = json.Unmarshal(content, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
