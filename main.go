package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	h "github.com/jedruniu/healthchecker/healthchecker"
)

func main() {
	ctx := context.Background()

	fileCheck := h.HealthCheck{
		Name:            "file based one",
		FailedThreshold: 10,
		PassedThreshold: 3,
		Interval:        1 * time.Second,
		S:               h.NewFileBased("testFile.txt"),
	}

	apiCheck := h.HealthCheck{
		Name:            "google endpoint",
		FailedThreshold: 10,
		PassedThreshold: 3,
		Interval:        2 * time.Second,
		S:               h.NewApiCallBased("http://google.com"),
	}

	redisCheck := h.HealthCheck{
		Name:            "get some key from redis",
		FailedThreshold: 1,
		PassedThreshold: 1,
		Interval:        3 * time.Second,
		S:               h.NewRedisBased("some_key"),
	}

	shellCheck := h.HealthCheck{
		Name:            "based on cmd",
		FailedThreshold: 1,
		PassedThreshold: 1,
		Interval:        2 * time.Second,
		S:               h.NewShellBased([]string{"true"}),
	}

	fileCheck.Run(ctx)
	apiCheck.Run(ctx)
	redisCheck.Run(ctx)
	shellCheck.Run(ctx)

	s := server{healths: []HealthReporter{&fileCheck, &apiCheck, &redisCheck, &shellCheck}}

	http.HandleFunc("/health", s.healthEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

type HealthReporter interface {
	IsHealthy() bool
}

type server struct {
	healths []HealthReporter
}

func (s server) healthEndpoint(w http.ResponseWriter, r *http.Request) {
	var content string
	for _, health := range s.healths {
		singleHealth := fmt.Sprintln(health, health.IsHealthy())
		content += singleHealth
	}
	w.Write([]byte(content))
}
