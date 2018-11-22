package healthchecker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type HealthReporter interface {
	IsHealthy() bool
}

type RunReporter interface {
	HealthReporter
	Runner
}

func HealthChecksFromConfig(cfg *[]Config) []RunReporter {
	var healthChecks []RunReporter

	for _, c := range *cfg {
		// create common part of Healthchecker
		check := &HealthCheck{
			Name:            c.Name,
			FailedThreshold: c.FailedThreshold,
			PassedThreshold: c.PassedThreshold,
			Interval:        time.Duration(c.Interval) * time.Second,
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
