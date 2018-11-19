package main

import (
	"fmt"
	"time"
)

func main() {
	healthcheck := fileBasedHealthCheck{filename: "testFile.txt", interval: 5 * time.Second}

	for {
		fmt.Println(healthcheck.isHealthy())
		time.Sleep(healthcheck.getInterval())
	}
}
