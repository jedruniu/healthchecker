package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	h "github.com/jedruniu/healthchecker/healthchecker"
)

func main() {
	ctx := context.Background()

	fileCheck := h.HealthCheck{
		Name: "file based one",
		FailedThreshold: 10,
		PassedThreshold: 3,
		Interval: 1*time.Second,
		S: h.NewFileBased("testFile.txt"),
	}

	apiCheck := h.HealthCheck{
		Name: "google endpoint",
		FailedThreshold: 10,
		PassedThreshold: 3,
		Interval: 2*time.Second,
		S: h.NewApiCallBased("http://google.com"),
	}

	redisCheck := h.HealthCheck{
		Name: "get some key from redis",
		FailedThreshold: 1,
		PassedThreshold: 1,
		Interval: 3*time.Second,
		S: h.NewRedisBased("some_key"),
	}

	fileCheck.Run(ctx)
	apiCheck.Run(ctx)
	redisCheck.Run(ctx)

	// TODO implement server to fetch data
	// TODO implement bash script checks

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
