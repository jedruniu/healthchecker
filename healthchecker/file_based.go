package healthchecker

import (
	"fmt"
	"os"
	"time"
)

type fileBasedHealthCheck struct {
	filename string
	interval time.Duration
}

// returns true if file was touched in less than minute, false otherwise
func (fbhc *fileBasedHealthCheck) isHealthy() bool {
	fileInfo, err := os.Stat(fbhc.filename)
	if err != nil {
		fmt.Printf("could not get info for file %s, err %v\n", fbhc.filename, err)
		return false
	}
	return (time.Now().Sub(fileInfo.ModTime()) < 1*time.Minute)
}

func (fbhc *fileBasedHealthCheck) getInterval() time.Duration {
	return fbhc.interval
}
