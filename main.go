package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jedruniu/healthcheck/healthcheck"
)

var port int
var configPath string

func init() {
	flag.IntVar(&port, "port", 8080, "server port on which health information will be accessible")
	flag.StringVar(&configPath, "config", "config.json", "config path")
}

func main() {
	flag.Parse()

	cfg, err := healthcheck.ReadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	healthChecks := healthcheck.HealthChecksFromConfig(cfg)

	ctx := context.Background()
	for _, hc := range healthChecks {
		hc.Run(ctx)
	}

	s := healthcheck.Server{Healths: healthChecks}
	http.HandleFunc("/health", s.HealthEndpoint)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
