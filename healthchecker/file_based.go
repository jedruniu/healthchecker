package healthchecker

import (
	"fmt"
	"os"
	"time"
)

type fileBasedHealthCheck struct {
	filename string
	*CommonHealthCheck
}

func NewFileBasedHealthCheck(filename string, interval time.Duration) HealthChecker {
	return &fileBasedHealthCheck{filename, &CommonHealthCheck{interval}}
}

// returns true if file was touched in less than minute, false otherwise
func (hc *fileBasedHealthCheck) IsHealthy() bool {
	fileInfo, err := os.Stat(hc.filename)
	if err != nil {
		fmt.Printf("could not get info for file %s, err %v\n", hc.filename, err)
		return false
	}
	return (time.Now().Sub(fileInfo.ModTime()) < 1*time.Minute)
}
